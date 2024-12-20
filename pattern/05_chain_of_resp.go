package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
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

