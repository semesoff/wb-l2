Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Будут выведены значения: 1, 2, 3, 4, 5, 6, 7, 8 в случайном порядке и бесконечно кол-во 0.

Объяснение:
Вывод бесконечного кол-ва 0, можно объяснить тем, что горутина merge читает данные из каналов и когда там нет уже данных, то мы получаем nil значение, а то есть 0. Для исправления этой проблемы нужно добавить проверку на то, что канал закрыт.
```
