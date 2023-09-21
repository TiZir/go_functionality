package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	// Отправляем значения в канал
	go func() {
		for {
			data := time.Now().Format("2006-01-02 15:04:05")
			ch <- data
			time.Sleep(time.Second)
		}
	}()

	// Таймер
	go func() {
		defer wg.Done()
		for {
			select {
			case <-time.After(time.Second * 12):
				close(ch)
				fmt.Println("Time is over. Closing channel...")
				return
			}
		}
	}()

	// Читаем значения из канала
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					fmt.Println("Канал закрыт")
					return
				}
				fmt.Println("Прочитано значение из канала:", val)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Program is finished")

}
