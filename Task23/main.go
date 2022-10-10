package main

import "fmt"

// Простенькая функция для удаления
func Delete[T any](slice []T, deletElement int) []T {
	// Создаем срез на основе начально до удаеямого элемента, и добовляем оставшие без удаляемого id
	return append(slice[:deletElement], slice[deletElement+1:]...)
}

// Удлание элемента но создает неупорядочный массив (скорее всего быстрее так как всего 1 раз пересоздает массив)
func DeleteNoOrder(slice []int, i int) []int {
	//Очень быстро
	last := len(slice) - 1
	slice[i] = slice[last]
	return slice[:last]
}

func main() {
	// Даннные срез
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	slice2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// До удаления, удаление, после удаления
	fmt.Println(slice)
	fmt.Println(Delete(slice, 4))

	// Прмиер 2 функиции
	fmt.Println(slice2)
	fmt.Println(DeleteNoOrder(slice2, 4))

}
