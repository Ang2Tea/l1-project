package main

import "fmt"

// Функция ввода чисел из последовательности
func GetNums(ch chan int, a ...int) {
	for _, val := range a {
		ch <- val
	}
	close(ch)
}

// Произведение чисел и 2
func Squares(nums, result chan int) {
	for val := range nums {
		result <- val * 2
	}
	close(result)
}

func main() {
	// Каналы для работы
	naturals := make(chan int)
	squares := make(chan int)

	// Запуск горутин
	go GetNums(naturals, 2, 4, 6, 8, 10)
	go Squares(naturals, squares)

	// Вывод
	for val := range squares {
		fmt.Printf("%d ", val)
	}
}
