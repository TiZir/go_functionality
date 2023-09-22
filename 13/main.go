package main

import "fmt"

func main() {
	a, b := 4, 5
	fmt.Println(a, b) // до смены
	b, a = a, b
	fmt.Println(a, b) // после

}
