package main

import "fmt"

func main() {
	// Словарь групп значений
	order := make(map[int][]float64)

	// Значения
	a := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 19, 21}

	// Рапределение по значения
	for _, value := range a {
		// Преобразуем число в десятковую группу
		group := int(value/10) * 10

		// Добовляем значение
		order[group] = append(order[group], value)
	}

	// Выводим
	for key, value := range order {
		fmt.Printf("%d: %+v\n", key, value)
	}
}
