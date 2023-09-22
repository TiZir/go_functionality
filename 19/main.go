package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse_string(s string) string {
	runes := []rune(s)
	length := len(runes)
	// Идес слева и справа по строке
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	str := "главрыба"

	// Проверяем строку на наличие символов unicode
	if !utf8.ValidString(str) {
		fmt.Println("Строка содержит некорректные символы")
		return
	}

	reversed := reverse_string(str)
	fmt.Println(reversed)
}
