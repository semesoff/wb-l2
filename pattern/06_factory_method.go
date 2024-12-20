package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Плюсы
1. Легко добавлять новые классы продуктов, не изменяя существующий код.
2. Логика создания объектов сосредоточена в одном месте.

Минусы
1. Добавление новых классов и методов может усложнить код.
2. Для каждого нового типа продукта необходимо создавать новый подкласс.
*/

// Product интерфейс продукта
type Product interface {
	Use() string
}

// ConcreteProductA конкретный продукт A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

// ConcreteProductB конкретный продукт B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}

// Creator интерфейс создателя
type Creator interface {
	FactoryMethod() Product
}

// ConcreteCreatorA конкретный создатель A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() Product {
	return &ConcreteProductA{}
}

// ConcreteCreatorB конкретный создатель B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

// Client code
// func main() {
// 	var creator Creator

// 	creator = &ConcreteCreatorA{}
// 	productA := creator.FactoryMethod()
// 	fmt.Println(productA.Use())

// 	creator = &ConcreteCreatorB{}
// 	productB := creator.FactoryMethod()
// 	fmt.Println(productB.Use())
// }
