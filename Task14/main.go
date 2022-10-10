package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

// Узнал о сбособах https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go

func main() {
	// Добавление значений в словарь для более удобой работы
	mass := make(map[string]interface{})
	mass["var1"] = "Сорок два"
	mass["var2"] = 42
	mass["var3"] = 42.42
	mass["var4"] = true
	mass["var5"] = make(chan int)
	mass["var6"] = [...]string{"42", "Сорок два", "Quadraginta duo"}
	mass["var7"] = Person{"Це Ганя"}

	// Форматирую что бы узнать тип
	fmt.Println("Опредедение через форматирование строки")

	for key, value := range mass {
		fmt.Printf("%s = %T\n", key, value)
	}
	// С помощью TypeOf(возращает перечесение типа Type)
	fmt.Println("\nОпредедение функцию TypeOf")
	for key, value := range mass {
		fmt.Printf("%s = %+v\n", key, reflect.TypeOf(value))
	}
	// С помощью ValueOf(возращает перечесение типа Kind)
	fmt.Println("\nОпредедение функцию ValueOf")
	for key, value := range mass {
		fmt.Printf("%s = %+v\n", key, reflect.ValueOf(value).Kind())
	}

}
