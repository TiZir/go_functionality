package main

import (
	"bytes"
	"fmt"
)

var justString string

// Cоздаем буфер buf типа bytes.Buffer и записываем в него огромную строку
func some_func() {
	var buf bytes.Buffer
	buf.WriteString(create_huge(1 << 10))
	justString = buf.String()[:100]
}

func main() {
	some_func()
	fmt.Println(justString)
}

// Используем буфер buf типа bytes.Buffer, чтобы создать огромную строку, добавляя по одному символу в буфер на каждой итерации цикла.
func create_huge(size int) string {
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteByte('a')
	}
	return buf.String()
}
