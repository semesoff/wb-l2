package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// cd - функция перехода по директориям
func cd(args []string) error {
	// если меньше двух аргументов, то возвращаем ошибку
	if len(args) < 2 {
		return fmt.Errorf("cd: missing argument")
	}
	// меняем текущий рабочий каталог на указанный путь
	return os.Chdir(args[1])
}

// pwd - функция вывода текущего рабочего каталога
func pwd() (string, error) {
	return os.Getwd()
}

// echo - функция вывода аргументов
func echo(args []string) string {
	// соединяем все аргументы в одну строку через пробел
	return strings.Join(args[1:], " ")
}

// kill - функция завершения процесса
func kill(args []string) error {
	// если меньше двух аргументов, то возвращаем ошибку
	if len(args) < 2 {
		return fmt.Errorf("kill: missing argument")
	}
	// завершаем процесс по указанному идентификатору
	cmd := exec.Command("kill", args[1])
	return cmd.Run()
}

// ps - функция вывода списка процессов
func ps() (string, error) {
	// запускаем команду ps
	cmd := exec.Command("ps")
	// получаем вывод команды
	output, err := cmd.Output()
	return string(output), err
}

// функция выполнения команды
func executeCommand(command string) {
	// разбиваем строку на аргументы по пробелам
	args := strings.Fields(command)
	if len(args) == 0 {
		return
	}

	// проверяем первый аргумент
	switch args[0] {
	case "cd":
		if err := cd(args); err != nil {
			fmt.Println(err)
		}
	case "pwd":
		if dir, err := pwd(); err == nil {
			fmt.Println(dir)
		} else {
			fmt.Println(err)
		}
	case "echo":
		fmt.Println(echo(args))
	case "kill":
		if err := kill(args); err != nil {
			fmt.Println(err)
		}
	case "ps":
		if output, err := ps(); err == nil {
			fmt.Println(output)
		} else {
			fmt.Println(err)
		}
	default:
		// запускаем команду, если она не встроенная
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout // перенаправляем вывод команды в стандартный вывод
		cmd.Stderr = os.Stderr // перенаправляем вывод ошибок команды в стандартный вывод ошибок
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "\\quit" {
			break
		}
		executeCommand(input)
	}
}
