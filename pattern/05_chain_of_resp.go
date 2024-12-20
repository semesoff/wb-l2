package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
События (например, нажатие кнопки) передаются по цепочке обработчиков, пока один из них не обработает событие.

Плюсы
1. Ослабление связи, отправитель запроса не знает, какой объект в цепочке обработает запрос.
2. Можно легко изменять цепочку обработчиков.
3. аждый обработчик отвечает только за свою часть работы.

Минусы
1. Трудно отследить, какой обработчик в цепочке обработал запрос.
*/

// Handler интерфейс обработчика
type Handler interface {
	SetNext(handler Handler)
	Handle(request string)
}

// BaseHandler базовый обработчик, реализующий общую логику установки следующего обработчика
type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(request string) {
	if h.next != nil {
		h.next.Handle(request)
	}
}

// ConcreteHandlerA конкретный обработчик A
type ConcreteHandlerA struct {
	BaseHandler
}

func (h *ConcreteHandlerA) Handle(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

// ConcreteHandlerB конкретный обработчик B
type ConcreteHandlerB struct {
	BaseHandler
}

func (h *ConcreteHandlerB) Handle(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

// ConcreteHandlerDefault обработчик по умолчанию
type ConcreteHandlerDefault struct {
	BaseHandler
}

// Handle обработчик по умолчанию, если не найден обработчик для данного
func (h *ConcreteHandlerDefault) Handle(request string) {
	fmt.Println("No one can handle the request")
}

// ConcreteHandlerC конкретный обработчик C
type ConcreteHandlerC struct {
	BaseHandler
}

func (h *ConcreteHandlerC) Handle(request string) {
	if request == "C" {
		fmt.Println("ConcreteHandlerC handled the request")
	} else {
		h.BaseHandler.Handle(request)
	}
}

// Client code
// func main() {
// 	handlerA := &ConcreteHandlerA{}
// 	handlerB := &ConcreteHandlerB{}
// 	handlerC := &ConcreteHandlerC{}
//  handlerDefault := &ConcreteHandlerDefault{}

// 	handlerA.SetNext(handlerB)
// 	handlerB.SetNext(handlerC)
//  handlerC.SetNext(handlerDefault)

// 	requests := []string{"A", "B", "C", "D"}

// 	for _, request := range requests {
// 		fmt.Printf("Handling request: %s\n", request)
// 		handlerA.Handle(request)
// 	}
// }
