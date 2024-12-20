package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн «Посетитель» позволяет уменьшить количество методов в классах,
переместив их в отдельные классы-посетители.

Плюсы
1.Упрощает добавление новых операций.
2.Объединяет родственные операции в одном месте.
3.Позволяет посещать объекты, не изменяя их классы.

Минусы
1.Усложняет структуру кода.
2.Может нарушить инкапсуляцию объектов.
3.Требует изменения всех классов, если добавляется новый тип элемента.
*/

// Element интерфейс, который принимает посетителя
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA конкретный элемент A
type ConcreteElementA struct {
	Value string
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// ConcreteElementB конкретный элемент B
type ConcreteElementB struct {
	Value int
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor конкретный посетитель, реализующий Visitor интерфейс
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("Processing ConcreteElementA with value: %s\n", element.Value)
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("Processing ConcreteElementB with value: %d\n", element.Value)
}

// func main() {
// 	elements := []Element{
// 		&ConcreteElementA{Value: "Hello"},
// 		&ConcreteElementB{Value: 42},
// 	}

// 	visitor := &ConcreteVisitor{}

// 	for _, element := range elements {
// 		element.Accept(visitor)
// 	}
// }
