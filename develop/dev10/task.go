package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

// опции
type Options struct {
	timeout time.Duration
}

// парсинг флагов
func parseFlags() Options {
	var options Options
	flag.DurationVar(&options.timeout, "timeout", 10*time.Second, "connection timeout")
	flag.Parse()
	return options
}

func telnet(address string, options Options) {
	// установка соединения
	conn, err := net.DialTimeout("tcp", address, options.timeout)
	if err != nil {
		fmt.Println("Error dialing:", err)
		return
	}
	// закрытие соединения
	defer conn.Close()

	done := make(chan struct{})
	sigChan := make(chan os.Signal, 1)
	// обработка сигнала прерывания
	// SIGTERM - kill, SIGINT - ctrl+c
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// копирование данных из stdin в сокет
	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			fmt.Println("Error copying:", err)
		}
		// передача сигнала о завершении работы
		done <- struct{}{}
	}()

	// копирование данных из сокета в stdout
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			fmt.Println("Error copying:", err)
		}
		// передача сигнала о завершении работы
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-sigChan:
		fmt.Println("Signal received, shutting down...")
	}
}

func main() {
	// парсинг флагов
	options := parseFlags()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		return
	}

	host := flag.Arg(0)                         // получаем первый аргумент
	port := flag.Arg(1)                         // получаем второй аргумент
	address := fmt.Sprintf("%s:%s", host, port) // "host:port"

	telnet(address, options)
}
