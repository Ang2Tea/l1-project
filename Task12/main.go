package main

import "fmt"

func main() {
	// Данные
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	// Определение уникальных значений
	variety := make(map[string]int)
	for _, el := range data {
		variety[el] = 1
	}
	// Преоброзование их в массив
	var result []string
	for key := range variety {
		result = append(result, key)
	}
	fmt.Println(result)
}
