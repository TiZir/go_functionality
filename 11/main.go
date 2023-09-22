package main

import "fmt"

func main() {
	// задаем два неупорядоченных множества
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// создаем карту для хранения элементов первого множества
	set1Map := make(map[int]bool)

	// добавляем элементы первого множества в карту
	for _, element := range set1 {
		set1Map[element] = true
	}

	// создаем слайс для хранения пересечения множеств
	intersection := []int{}

	// проходим по элементам второго множества и проверяем, есть ли они в первом множестве
	for _, element := range set2 {
		if set1Map[element] {
			intersection = append(intersection, element)
		}
	}

	// выводим результаты
	fmt.Println("Пересечение множеств:", intersection)
}
