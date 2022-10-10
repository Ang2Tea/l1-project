package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создание больших чисел
	// Дальше все десятки в конце обозначает что нам нужно вернуть значение в десятичной форме
	var num1, _ = new(big.Int).SetString("123456782818312731623", 10)
	var num2, _ = new(big.Int).SetString("123456782818312731623", 10)
	// Переменная для результата
	result := big.NewInt(1)
	// Опреации с числами
	fmt.Println(result.Add(num1, num2).Text(10))
	fmt.Println(result.Sub(num1, num2).Text(10))
	fmt.Println(result.Mul(num1, num2).Text(10))
	fmt.Println(result.Div(num1, num2).Text(10))
}
