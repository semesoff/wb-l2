package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func getTime(ntpServer string) (time.Time, error) {
	currentTime, err := ntp.Time(ntpServer)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime, nil
}

func main() {
	currentTime, err := getTime("time.google.com")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time: %s\n", currentTime)
}
