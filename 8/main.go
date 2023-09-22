package main

import "fmt"

func main() {
	var num int64 = 42
	i := 2 // номер бита, который нужно установить

	// Установить i-й бит в 1
	num |= 1 << i
	fmt.Printf("Установлен i-й бит в 1: %d\n", num)

	// Установить i-й бит в 0
	num &= ^(1 << i)
	fmt.Printf("Установлен i-й бит в 0: %d\n", num)
}
