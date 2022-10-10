package main

import (
	"fmt"
)

// Переписал с https://github.com/Ang2Tea/quick_sort_rust
func main() {
	// Исходдный массив
	data := []int{7, 1, 5, 6, 2, 2}
	fmt.Println(data)

	// Быстрая сортировка
	QuickSort(&data)
	// На выходе
	fmt.Println(data)
}

func QuickSort(data *[]int) {
	// Если массив меньше 2 элементов функция начент сварачиваться
	if len(*data) < 2 {
		return
	}
	// Елемен с которым будут сравиваться другие для распределения
	zeroElem := (*data)[0]

	// Распределение
	min, max := []int{}, []int{}
	for _, value := range (*data)[1:] {
		if value < zeroElem {
			min = append(min, value)
		}
		if value >= zeroElem {
			max = append(max, value)
		}
	}
	// Сортировка разделенных частей
	QuickSort(&min)
	QuickSort(&max)

	// Вставка в исходный массив новых данных
	min = append(min, zeroElem)
	min = append(min, max...)
	copy(*data, min)
}
