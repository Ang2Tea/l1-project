package main

import "fmt"

func main() {
	// Множества
	a1 := []int{0, 4, 8, 12, 16}
	a2 := []int{0, 2, 4, 6, 8}
	// Соедение множеств для удобства
	a1 = append(a1, a2...)

	// Map для подчета ценности
	intersection := make(map[int]int)
	for _, el := range a1 {
		// Если элемен встречается больше 1 раза ему увеличивается соимость
		intersection[el] += 1
	}

	// Смотрим на ценнсть переменной и так где больше 1 выводим
	var result []int
	for key, value := range intersection {
		if value > 1 {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}
