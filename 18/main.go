package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	// Запускаем гоурутины с увеличение счетчика
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(count) //500 от каждой Гоурутины
}
