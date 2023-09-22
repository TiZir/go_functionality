package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10} // создаем слайс с  данными
	wg := sync.WaitGroup{}           // создаем WaitGroup для ожидания завершения всех горутин
	sum := 0                         // счетчик для суммы

	// Запускаем горутины для обработки данных
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			squared := n * n // квадрат
			sum += squared   // сумма квадратов
		}(num)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("Сумма квадратов:", sum)
}
