package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Плюсы
1.Позволяет легко реализовать отмену и повтор операций.
2.Разделяет объекты, инициирующие операции, и объекты, которые их выполняют.

Минусы
1. Усложняет код за счет введения большого количества дополнительных классов.
*/

// Command интерфейс команды
type Command interface {
	Execute()
}

// LightReceiver получатель команды
type LightReceiver struct{}

func (l *LightReceiver) On() {
	fmt.Println("Light is On")
}

func (l *LightReceiver) Off() {
	fmt.Println("Light is Off")
}

// LightOnCommand команда включения света
type LightOnCommand struct {
	light *LightReceiver
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// LightOffCommand команда выключения света
type LightOffCommand struct {
	light *LightReceiver
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// RemoteControlInvoker инициатор команды
type RemoteControlInvoker struct {
	command Command
}

func (r *RemoteControlInvoker) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControlInvoker) PressButton() {
	r.command.Execute()
}

// Client code
// func main() {
// 	light := &LightReceiver{}

// 	lightOn := &LightOnCommand{light: light}
// 	lightOff := &LightOffCommand{light: light}

// 	remote := &RemoteControlInvoker{}

// 	remote.SetCommand(lightOn)
// 	remote.PressButton()

// 	remote.SetCommand(lightOff)
// 	remote.PressButton()
// }
