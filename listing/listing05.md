Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод:
error

Объяснение:
Мы объявляем переменную err типа error.
error - это интерфейс, который определяет один метод Error() string. Переменной err присваиваем значение функции test, которая возвращает указатель на customError, равный nil. У структуры customError есть метод Error() string, значит структура реализует интерфейс error. 
Т.к. переменная err имеет тип, то она не вляется пустым интерфейсом, поэтому условие err != nil будет true, и в итоге напечатается error.
```
