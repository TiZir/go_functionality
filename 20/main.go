package main

import (
	"fmt"
	"strings"
)

func reverse_words(s string) string {
	words := strings.Fields(s)
	reversed := make([]string, len(words))
	// Меняем местами символ от начала и от конца местами
	for i := len(words) - 1; i >= 0; i-- {
		reversed[len(words)-1-i] = words[i]
	}

	return strings.Join(reversed, " ")
}

func main() {
	str := "snow dog sun"
	reversed := reverse_words(str)
	fmt.Println(reversed)
}
