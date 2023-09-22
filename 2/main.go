package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10} // создаем slice
	var wg sync.WaitGroup         // создаем WaitGroup для ожидания завершения всех горутин
	for _, num := range nums {
		wg.Add(1)        // добавляем горутину в WaitGroup
		go func(x int) { // запускаем горутину
			defer wg.Done()    // уменьшаем счетчик горутин в WaitGroup при завершении
			fmt.Println(x * x) // выводим квадрат числа в stdout
		}(num)
	}
	wg.Wait() // ждем завершения всех горутин
}
