package main

import (
	"fmt"
)

// Функция для подчета квалратов
func NumSquareAsync(result chan<- int, nums ...int) {
	// Закрытие горотуны в конце выполенения функции
	defer close(result)

	for _, val := range nums {
		// Отправка данных в горутину
		result <- val * val
	}
}

func main() {
	// Создание канала для передачи данных из горотины
	result := make(chan int)

	// Запуск горутины
	go NumSquareAsync(result, 2, 4, 6, 8, 10)

	// Вывод значений из горутины пока она не закроется
	for val := range result {
		fmt.Printf("%d ", val)
	}
}
