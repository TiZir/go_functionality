package main

import (
	"fmt"
)

func check_type(i interface{}) {
	switch s := i.(type) {
	case int:
		fmt.Println("Это int", s)
	case string:
		fmt.Println("Это string", s)
	case bool:
		fmt.Println("Это bool", s)
	case chan int:
		fmt.Println("Это chan int", s)
	default:
	}
}

func main() {
	var i interface{}

	// присваиваемзначение разных типов
	i = 12
	check_type(i)

	i = "cat"
	check_type(i)

	i = true
	check_type(i)

	i = make(chan int)
	check_type(i)
}
