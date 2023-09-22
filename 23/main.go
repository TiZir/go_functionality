package main

import "fmt"

func remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	i := 1           // индекс удаления
	fmt.Println(arr) // до
	arr = remove(arr, i)
	fmt.Println(arr) // после

}
