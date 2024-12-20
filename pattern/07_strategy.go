package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Необходимо определить семейство алгоритмов, инкапсулировать каждый из них и сделать их взаимозаменяемыми.
Требуется выбрать алгоритм или поведение во время выполнения.

Плюсы
1. Легко менять алгоритмы или поведения объекта во время выполнения.
2. Алгоритмы изолированы в собственных классах, что упрощает поддержку и расширение.
3. Легко добавлять новые стратегии без изменения существующего кода.
4. Позволяет избежать множественных условных операторов.

Минусы
1. Увеличение количества классов и объектов.
*/

// Strategy интерфейс стратегии
type Strategy interface {
	Algorithm()
}

// Context контекст, использующий стратегию
type Context struct {
	strategy Strategy
}

// Изменение алгоритма
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Выполнение алгоритма
func (c *Context) Operation() {
	c.strategy.Algorithm()
}

// ConcreteStrategyA конкретная стратегия A
type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) Algorithm() {
	fmt.Println("Executing Algorithm A")
}

// ConcreteStrategyB конкретная стратегия B
type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) Algorithm() {
	fmt.Println("Executing Algorithm B")
}

// Client code
// func main() {
// 	context := &Context{}

// 	strategyA := &ConcreteStrategyA{}
// 	strategyB := &ConcreteStrategyB{}

// 	context.SetStrategy(strategyA)
// 	context.Operation()

// 	context.SetStrategy(strategyB)
// 	context.Operation()
// }
