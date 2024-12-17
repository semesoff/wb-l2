package pattern

import "fmt"

/*
   Реализовать паттерн «фасад».
   Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
   https://en.wikipedia.org/wiki/Facade_pattern
*/

// Плюсы
// 1. Изолирует клиентов от компонентов сложной системы.
// 2. Предоставляет простой интерфейс для работы со сложной системой.
// 3. Изменения в сложной системе не влияют на клиентов, если интерфйес остается неизменным.
// Минусы
// 1. Фасад может не предоставлять весь функционал сложной системы.

// Database представляет сложную подсистему
type Database struct{}

func (d *Database) Start() {
	fmt.Println("Database started.")
}

// Server представляет другую сложную подсистему
type Server struct{}

func (s *Server) Start() {
	fmt.Println("Server started.")
}

// Facade предоставляет простой интерфейс для взаимодействия с подсистемами
type Facade struct {
	db     *Database
	server *Server
}

// NewFacade создает новый экземпляр фасада
func NewFacade() *Facade {
	return &Facade{
		db:     &Database{},
		server: &Server{},
	}
}

// Start запускает все подсистемы через фасад
func (f *Facade) Start() {
	f.db.Start()
	f.server.Start()
}

// func main() {
// 	facade := NewFacade()
// 	facade.Start()
// }
