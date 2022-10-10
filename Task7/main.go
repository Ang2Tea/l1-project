package main

import (
	"fmt"
	"sync"
)

// Структура своего Map
type MyMap struct {
	mutex sync.RWMutex
	data  map[string]string
}

// Иницилизация новой структуры
func New() *MyMap {
	mm := MyMap{
		data: make(map[string]string),
	}

	return &mm
}

// Запись данных в map
func (mm *MyMap) Write(key, value string) {
	defer mm.mutex.Unlock()
	mm.mutex.Lock()
	mm.data[key] = value
}

// Чтнение данных из канала
func (mm *MyMap) Read(key string, result chan<- string) {
	defer mm.mutex.RUnlock()
	mm.mutex.RLock()
	result <- mm.data[key]
}

// Декорация записи для того что бы ожидать записи всех данных
func TestW(wg *sync.WaitGroup, mm *MyMap, key, value string) {
	defer wg.Done()
	mm.Write(key, value)
}

func main() {
	// Канал для результато
	result := make(chan string, 5)

	// Создание Map
	maP := New()
	// Создание wg для ожидания завершение записи
	wg1 := sync.WaitGroup{}
	wg1.Add(5)

	go TestW(&wg1, maP, "1", "Forty two")
	go TestW(&wg1, maP, "2", "Сорок два")
	go TestW(&wg1, maP, "3", "Quarante-deux")
	go TestW(&wg1, maP, "4", "Zweiundvierzig")
	go TestW(&wg1, maP, "5", "Quarantadue")
	wg1.Wait()

	// Начала чтение
	go maP.Read("5", result)
	go maP.Read("3", result)
	go maP.Read("4", result)
	go maP.Read("1", result)
	go maP.Read("2", result)

	// После завершение записи падает ошибка, но я не знаю как ее пофиксить
	for value := range result {
		fmt.Println(value)
	}
}
