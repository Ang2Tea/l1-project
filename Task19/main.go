package main

import "fmt"

func Reverse(str string) string {
	// Длина массива
	length := len(str)
	// Создание массива сиволов
	chars := make([]rune, length)
	// Изменение местами значений
	for _, value := range str {
		length--
		chars[length] = value
	}
	// Возращение строк
	return string(chars)
}

func main() {
	// Строки для примера
	strs := [...]string{
		"Сорок два",
		"Эй, у меня, может, и самая громкая глотка на Столичной пустоши, но ещё — у меня самые большие уши.",
		"абырвалг",
	}

	// Вывод перевернутых строк
	for _, value := range strs {
		fmt.Println(Reverse(value))
	}
}
