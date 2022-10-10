package main

import "fmt"

// Фукнция ввода чисел в горутину
func GetNums(numbers chan<- int, nums ...int) {
	for _, val := range nums {
		numbers <- val
	}
	close(numbers)
}

// Функция получения квадрата горутины
func Square(squares chan<- int, numbers <-chan int) {
	for val := range numbers {
		squares <- val * val
	}
	close(squares)
}

// Вычисление частей суммы в разных потоках
func CountSumParts(squares <-chan int, partSum chan<- int) {
	presum := 0
	for value := range squares {
		presum += value
	}
	partSum <- presum
}

// Чтение частей суммы и вычеленеие итого
func Result(part <-chan int) int {
	var sum int
	for i := 0; i < cap(part); i++ {
		sum += <-part
	}
	return sum
}

func main() {
	// Создание горутин
	numbers := make(chan int)
	squares := make(chan int)

	// Отправка чисел в отдельном горутины
	go GetNums(numbers, 2, 4, 6, 8, 10)

	// Начала подчета квадратов
	go Square(squares, numbers)

	// Количество потоков для обработки суммы
	const countThreads = 4
	partSum := make(chan int, countThreads)
	defer close(partSum)
	for i := 0; i < countThreads; i++ {
		go CountSumParts(squares, partSum)
	}

	// Чтение честей суммы и вычеленеие итого
	fmt.Println(Result(partSum))
}
