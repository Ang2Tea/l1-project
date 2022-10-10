package main

import (
	"fmt"
	"sync"
)

// Тип счетчика для конкурентной записи
type Counter struct {
	mutex sync.Mutex
	value int
}

// Метод для безопасной увелечения значения
func (c *Counter) Raise() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.value++
}

// Обертка над увелечением, что бы дожидаться окочания задачи
func WaitFunc(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	c.Raise()
}

func main() {
	//Обьект типа счетчик
	c := Counter{value: 0}
	// Создание WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	iter := 6
	wg.Add(iter)
	for i := 0; i < iter; i++ {
		// Запуск горутин
		go WaitFunc(&c, &wg)
	}

	// Ожидание и вывод ответа
	wg.Wait()
	fmt.Printf("Счет %d", c.value)
}
