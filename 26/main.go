package main

import (
	"fmt"
	"strings"
)

func is_unique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	// Создаем карту для хранения символов
	charMap := make(map[rune]bool)

	// Проходим по строке и проверяем, что каждый символ уникален
	for _, char := range str {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	fmt.Println(is_unique("abcd"))      // true
	fmt.Println(is_unique("abCdefAaf")) // false
	fmt.Println(is_unique("aabcd"))     // false
}
