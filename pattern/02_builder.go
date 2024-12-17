package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
// Плюсы
// Позволяет создавать объект поэтапно, вызывая методы строителя.
// Улучшает читаемость кода, так как каждый метод строителя отвечает за установку одного параметра.
// Минусы
// Может затруднить внедрение зависимостей.
// Увеличивает сложность кода из-за необходимости создания дополнительных классов и интерфейсов.
// Может быть избыточным для простых объектов, где достаточно использовать конструктор с параметрами.

// House представляет сложный объект, который будем создавать
type House struct {
	walls        string
	roof         string
	floors       int
	garage       bool
	swimmingPool bool
}

// HouseBuilder определяет интерфейс строителя
type HouseBuilder interface {
	SetWalls(walls string) HouseBuilder
	SetRoof(roof string) HouseBuilder
	SetFloors(floors int) HouseBuilder
	SetGarage(garage bool) HouseBuilder
	SetSwimmingPool(swimmingPool bool) HouseBuilder
	Build() House
}

// ConcreteHouseBuilder реализует интерфейс HouseBuilder
type ConcreteHouseBuilder struct {
	house House
}

func NewHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{}
}

func (b *ConcreteHouseBuilder) SetWalls(walls string) HouseBuilder {
	b.house.walls = walls
	return b
}

func (b *ConcreteHouseBuilder) SetRoof(roof string) HouseBuilder {
	b.house.roof = roof
	return b
}

func (b *ConcreteHouseBuilder) SetFloors(floors int) HouseBuilder {
	b.house.floors = floors
	return b
}

func (b *ConcreteHouseBuilder) SetGarage(garage bool) HouseBuilder {
	b.house.garage = garage
	return b
}

func (b *ConcreteHouseBuilder) SetSwimmingPool(swimmingPool bool) HouseBuilder {
	b.house.swimmingPool = swimmingPool
	return b
}

func (b *ConcreteHouseBuilder) Build() House {
	return b.house
}

// func main() {
// 	builder := NewHouseBuilder()
// 	house := builder.SetWalls("Brick").SetRoof("Tile").SetFloors(2).SetGarage(true).SetSwimmingPool(true).Build()
// 	fmt.Printf("House: %+v\n", house)
// }
