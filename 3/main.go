package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	sum := 0

	// Запускаем горутины для обработки данных
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			squared := n * n
			sum += squared
		}(num)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("Сумма квадратов:", sum)
}
