package main

import "fmt"

func main() {
	//Первый способ доступный не во всех языках
	first := 5
	second := 7
	fmt.Printf("Начальные значения: first=%d\tsecond=%d \n", first, second)

	first, second = second, first
	fmt.Printf("Первый способ: first=%d\tsecond=%d \n", first, second)

	//Второкой способ
	first = 5
	second = 7

	first += second         // first=13	second=7
	second = first - second // first=13	second=5
	first -= second         // first=7	second=5
	fmt.Printf("Второй способ: first=%d\tsecond=%d \n", first, second)

	//Третий способ
	first = 5
	second = 7

	first *= second         // first=35	second=7
	second = first / second // first=35	second=5
	first = first / second  // first=7	second=5
	fmt.Printf("Третий способ: first=%d\tsecond=%d \n", first, second)
}
