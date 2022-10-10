package main

import (
	"fmt"
)

func First() {
	fmt.Println("First")
}

func SercondAndThird(abort chan struct{}) {
	select {
	case <-abort:
		return
	default:
		fmt.Println("123")
	}
}

func main() {
	// Первый способ просто дождаться завершения горутины
	go First()

	// Второй способ создать канал в который будет при надобнасти закрытия будет посылаться данные и которые его уже будут закрыват
	abort1 := make(chan struct{})
	go SercondAndThird(abort1)
	abort1 <- struct{}{}

	// Третий способ похож со вторым, но теперь мы закрывает канал функцией close. Select в функции отробатывает так же
	abort2 := make(chan struct{})
	go SercondAndThird(abort2)
	close(abort2)
}
