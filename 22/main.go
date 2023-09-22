package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем два больших числа больше int64
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("12345678901234567890", 10)
	b.SetString("98765432109876543210", 10)

	// Делаем операция с числами
	add := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	// Выводим ответы
	fmt.Printf("a + b = %v\n", add)
	fmt.Printf("a - b = %v\n", sub)
	fmt.Printf("a * b = %v\n", mul)
	fmt.Printf("a / b = %v\n", div)
}
