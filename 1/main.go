package main

import "fmt"

// Определяется структура Human с двумя полями: name типа string и age типа int.
type Human struct {
	name string
	age  int
}

// Определяется метод SayHello для структуры Human. Метод SayHello выводит на экран приветствие с именем и возрастом человека.
func (h *Human) SayHello() {
	fmt.Printf("Привет, мое имя %s и мне %d лет.\n", h.name, h.age)
}

// Определяется структура Action, которая включает в себя структуру Human. Просто и с кратким обозначением в H
type Action struct {
	Human
	H Human
}

/*
В функции main создается переменная a типа Action, которая содержит структуру Human с именем "Бибер" и возрастом 18, а также
Долик с возрастом 18. Затем вызывается метод SayHello из структуры Human для переменной a и для переменной a.H,
который выводит на экран приветствие.
*/
func main() {
	a := Action{Human{"Бибер", 18}, Human{"Долик", 18}}
	a.SayHello()
	a.H.SayHello()
}
