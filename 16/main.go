package main

import (
	"fmt"
	"math/rand"
)

func quicksort(array []int) []int {

	if len(array) < 2 {
		return array
	} //базовый случай

	left, right := 0, len(array)-1   //левая и правая граница
	pivot := rand.Int() % len(array) //опорный

	array[pivot], array[right] = array[right], array[pivot] //опорный на конец для сравнения

	for i := range array {
		if array[i] < array[right] {
			array[i], array[left] = array[left], array[i]
			left++
		} //если меньше опорного, кидаем в левую часть и двигаемся на 1 элемент вправо
	}

	array[left], array[right] = array[right], array[left] //left pivot right

	quicksort(array[:left])   //сортируем left
	quicksort(array[left+1:]) //сортируем  right

	return array
}

func main() {
	mas := []int{2, 1, 3, 7, 5, 6, 4, 8, 9, 10, 11, -6}
	fmt.Printf("before\n") // вывод до
	for _, element := range mas {
		fmt.Printf("%d ", element)
	}
	fmt.Printf("\nafter\n") // вывод после
	mas = quicksort(mas)
	for _, element := range mas {
		fmt.Printf("%d ", element)
	}
}
