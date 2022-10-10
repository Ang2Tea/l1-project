package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

// Чтение консоли
func ReadConsole(texts chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		texts <- text
	}
}

// Функция читает канал и пишит их в консоль
func PrintText(texts <-chan string, close <-chan interface{}) {
	for {
		select {
		case mes := <-texts:
			fmt.Fprint(os.Stdout, mes)

		// Если канал close закрылся то он срабатывает и отрабатывает оставшие сообщения, после чего закрывает горутину
		case <-close:
			for mes := range texts {
				fmt.Fprint(os.Stdout, mes)
			}
			return
		}
	}
}

func main() {
	workTime := flag.Int("time", 10, "times")
	closeApp := make(chan interface{}, 1)
	textCh := make(chan string)

	go ReadConsole(textCh)
	go PrintText(textCh, closeApp)

	time.Sleep(time.Duration(*workTime) * time.Second)

	close(closeApp)
}
