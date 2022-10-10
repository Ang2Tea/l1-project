package main

import (
	"fmt"
	"time"
)

// Функция сна
func Sleep(seconds int) {
	// Ждет когда появится сообщение в канале
	// Который появится через second времени
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начала")

	Sleep(5)

	fmt.Println("Конец")
}
