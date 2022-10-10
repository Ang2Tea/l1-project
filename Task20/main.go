package main

import (
	"fmt"
	"strings"
)

// Фукнция перевертывания
func ReverseWord(str string) string {
	// Разбиение на слова
	// Если я правильно понял то если нужна настрока по разбиению сиволов можно использовать функцию Split
	words := strings.Fields(str)

	// Строка для результата
	length := len(words)
	newString := make([]string, length)

	// Перевертывание предложения
	for _, value := range words {
		length--
		newString[length] = value
	}

	return strings.Join(newString, " ")
}

func main() {
	// Строки для примера
	strs := [...]string{
		"Сорок два",
		"Эй, у меня, может, и самая громкая глотка на Столичной пустоши, но ещё — у меня самые большие уши",
		"абырвалг",
	}

	// Вывод перевернутых строк
	for _, value := range strs {
		fmt.Println(ReverseWord(value))
	}
}
