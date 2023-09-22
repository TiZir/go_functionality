package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
)

func binary_search(list [117]string, item string) (int, error) {
	left := 0           // начало файла
	right := 117        // конец файла
	for left <= right { // пока не дошли до конца
		mid := (left + right) / 2 // центр
		guess := list[mid]        // значение центра
		if guess == item {        // если угадали
			return mid, nil
		}
		if guess > item { // если вщяли больше
			right = mid - 1 // обрезаем все справа
		} else { // если взяли меньше
			left = mid + 1 // обрезаем все слева
		}

	}
	return 0, errors.New("can't found")
}

func main() {
	f, err := os.Open("17/names.txt") // открываем файл с данными
	if err != nil {
		panic(err)
	}
	defer f.Close()
	os.Stdin = f // отправляем данные в stdin
	var listNames [117]string 
	for i := 0; i < 117; i++ { // считываем данные из stdin
		if _, err := fmt.Scanf("%s", &listNames[i]); err != nil {
			fmt.Printf("Error: %v", err)
			return
		}
	}
	sort.Strings(listNames[:]) //сортируем список имен
	result, err := binary_search(listNames, "Mia") //ищем имя Mia
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("%d", result)
}
