package main

import "fmt"

// Метод бинарного поиска
func Search(num int, data []int) int {
	// Назначение минимального и максимальоо индекса
	first, last := 0, len(data)
	for first < last {
		// Нахождение среднего индекса
		middle := first + (last-1-first)/2
		// Если число по индексу подходит возращаем индекс
		if data[middle] == num {
			return middle
		}
		// Если нет то уменьшаем максимальный индекс или увеличиваем минимальный
		if data[middle] < num {
			first = middle + 1
		} else {
			last = middle - 1
		}
	}

	// Если ничего не нашли возращаем -1
	return -1
}

func main() {
	// Массив
	mass := []int{1, 3, 24, 10, 7, 15, 21, 28, 42, 45, 55}
	// Поиск
	index := Search(55, mass)
	// Вывод результата
	if index != -1 {
		fmt.Printf("mass[%d]=%d", index, mass[index])
	} else {
		fmt.Printf("Число не найдено")
	}

}
