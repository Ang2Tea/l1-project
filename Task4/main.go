package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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

// Чтение консоли
func ReadConsole(texts chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		texts <- text
	}
}

func main() {
	// Аргумент для считаывания воркетор при запуске
	workers := flag.Int("worker", 1, "workers")
	flag.Parse()
	// Канал сообщений
	textCh := make(chan string)
	// Канал для сигналов выхода
	quitApp := make(chan os.Signal, 1)

	// Канал для коретного работы вывода при завершение (нашел в интернете. Изначально пытался использовать default в select)
	closesGorotin := make(chan interface{}, 1)

	// Запуск чтения строки
	go ReadConsole(textCh)

	// Запуск обработчиков сообщений
	for i := 0; i < *workers; i++ {
		go PrintText(textCh, closesGorotin)
	}

	// Отправляет сообщене в кала закрытия если произошла одна из ошибок
	signal.Notify(quitApp, syscall.SIGTERM, syscall.SIGINT)
	// Ждет когда придет какой либо сигнал
	<-quitApp

	// Закрытие канало для корекного завершения работы
	close(textCh)
	close(closesGorotin)

}
