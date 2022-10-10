package main

import (
	"fmt"
	"strings"
)

// Функции проверки уникальности
func CheckUniq(str string) bool {
	// Перевод строки к одному регистру
	str = strings.ToLower(str)
	//Создание словоря для уник значений
	dict := make(map[rune]bool)
	// Цикл хождение по слову
	for _, value := range str {
		// Если значниея нет вернут дефолтное false
		// И пройдет в первое условие запишится в словарь
		if check := dict[value]; !check {
			dict[value] = true

		} else {
			// Если значение уже есть то будет false и он вывдет программу из функции вернуф false
			return false
		}
	}
	return true
}

func main() {
	strs := [...]string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}
	for _, value := range strs {
		fmt.Printf("%v ", CheckUniq(value))
	}
}
