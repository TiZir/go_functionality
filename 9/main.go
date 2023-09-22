package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	firstChan := make(chan int)
	secondChan := make(chan int)
	wg := new(sync.WaitGroup)

	wg.Add(3)

	// Гоурутина пишет данные в канал
	go func() {
		defer wg.Done()
		for _, val := range arr {
			firstChan <- val
		}
		close(firstChan)
	}()

	// Гоурутина умножает данные на 2
	go func() {
		defer wg.Done()
		for val := range firstChan {
			secondChan <- val * 2
		}
		close(secondChan)
	}()

	// Гоурутина выводит данные
	go func() {
		defer wg.Done()
		for val := range secondChan {
			fmt.Fprintln(os.Stdout, val)
		}
	}()

	wg.Wait()
}
