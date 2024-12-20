package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Поведение редактора зависит от режима (вставка, редактирование, выделение).

Плюсы
1. Избавляет от множества условных операторов, управляющих состоянием.
2. Каждый класс состояния отвечает только за свое поведение.

Минусы
1. Может привести к увеличению количества классов и объектов.
*/

// State интерфейс состояния
type State interface {
	Handle(manager *Manager)
}

// Manager контекст, использующий состояния
type Manager struct {
	state State
}

// SetState изменение состояния
func (c *Manager) SetState(state State) {
	c.state = state
}

// Request выполнение запроса
func (c *Manager) Request() {
	c.state.Handle(c)
}

// ConcreteStateA конкретное состояние A
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(manager *Manager) {
	fmt.Println("State A handling request.")
	manager.SetState(&ConcreteStateB{})
}

// ConcreteStateB конкретное состояние B
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(manager *Manager) {
	fmt.Println("State B handling request.")
	manager.SetState(&ConcreteStateA{})
}

// Client code
// func main() {
// 	context := &Context{state: &ConcreteStateA{}}

// 	context.Request()
// 	context.Request()
// 	context.Request()
// 	context.Request()
// }
