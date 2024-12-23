package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestCd(t *testing.T) {
	// Сохраняем текущую директорию
	originalDir, _ := os.Getwd()
	defer os.Chdir(originalDir)

	// Тестируем смену директории
	if err := cd([]string{"cd", "/tmp"}); err != nil {
		t.Errorf("cd failed: %v", err)
	}

	// Проверяем, что текущая директория изменилась
	newDir, _ := os.Getwd()
	expectedDir, _ := filepath.EvalSymlinks("/tmp")
	if newDir != expectedDir {
		t.Errorf("expected %s, got %s", expectedDir, newDir)
	}
}

func TestPwd(t *testing.T) {
	// Сохраняем текущую директорию
	originalDir, _ := os.Getwd()

	// Проверяем, что функция pwd возвращает текущую директорию
	if dir, err := pwd(); err != nil {
		t.Errorf("pwd failed: %v", err)
	} else if dir != originalDir {
		t.Errorf("expected %s, got %s", originalDir, dir)
	}
}

func TestEcho(t *testing.T) {
	// Проверяем, что функция echo корректно выводит аргументы
	result := echo([]string{"echo", "Hello,", "World!"})
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestKill(t *testing.T) {
	// Создаем процесс для тестирования
	cmd := exec.Command("sleep", "10")
	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start process: %v", err)
	}

	// Завершаем процесс
	if err := kill([]string{"kill", fmt.Sprintf("%d", cmd.Process.Pid)}); err != nil {
		t.Errorf("kill failed: %v", err)
	}

	// Проверяем, что процесс завершен
	if err := cmd.Wait(); err == nil {
		t.Errorf("expected process to be killed")
	}
}

func TestPs(t *testing.T) {
	// Проверяем, что функция ps возвращает список процессов
	if output, err := ps(); err != nil {
		t.Errorf("ps failed: %v", err)
	} else if !strings.Contains(output, "PID") {
		t.Errorf("expected output to contain PID, got %s", output)
	}
}

func TestExecuteCommand(t *testing.T) {
	// Тестируем выполнение команды echo
	output := captureOutput(func() {
		executeCommand("echo Hello, World!")
	})
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("expected %s, got %s", expected, output)
	}

	// Тестируем выполнение команды pwd
	originalDir, _ := os.Getwd()
	output = captureOutput(func() {
		executeCommand("pwd")
	})
	if strings.TrimSpace(output) != originalDir {
		t.Errorf("expected %s, got %s", originalDir, output)
	}
}

// captureOutput захватывает вывод функции
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout

	var buf strings.Builder
	io.Copy(&buf, r)
	return buf.String()
}
