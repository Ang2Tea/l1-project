package main

import (
	"fmt"
	"strings"
)

// Структура Human
type Human struct {
	Surname, Name string
}

// Метод устонавливающий значение фамилии
func (h *Human) SetSurname(Surname string) {
	h.Surname = Surname
}

// Метод устонавливающий значение имени
func (h *Human) SetName(surname string) {
	h.Surname = surname
}

// Структура Action
type Action struct {
	Human
	Age uint
}

// Переназначеный метод устонавливающий значение имени
func (a *Action) SetName(name string) {
	a.Name = strings.ToUpper(name)
}

// Не уноследованные метод
func (a *Action) SetAge(age uint) {
	a.Age = age
}

func main() {
	// Создания структур
	human := Human{Name: "Human name", Surname: "Human surname"}
	action := Action{
		Age:   10,
		Human: human}
	fmt.Printf("Начальное состояние\n %+v\n %+v\n", human, action)

	// Вызов унаследоного метода
	action.SetSurname("Action surname incapsulated")
	fmt.Printf("После использования унаследовоного метода, %+v\n", action)

	// Вызов собственного метода
	action.SetAge(300)
	fmt.Printf("После собственного метода, %+v\n", action)

	// Вызов переопределнной реалезации
	action.SetName("Action name")
	fmt.Printf("После использования переопределенного метода, %+v\n", action)

	// Вызов Базовой реалезации
	action.Human.SetName("Action name")
	fmt.Printf("После использования базового метода, %+v\n", action)

	fmt.Printf("Конечное состояние\n%+v\n%+v\n", human, action)
}
