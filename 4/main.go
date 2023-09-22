package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Создаем каналы для коммуникации между горутинами
	dataChan := make(chan string)   // Канал для записи данных
	doneChan := make(chan struct{}) // Канал для уведомления об окончании работы

	// Создаем и запускаем главную горутину, которая будет записывать данные в канал
	go func() {
		for {
			// Имитируем запись данных в канал
			data := time.Now().Format("2006-01-02 15:04:05")
			dataChan <- data
			// Задержка между записью данных
			time.Sleep(time.Second)
		}
	}()

	// Создаем и запускаем воркеров для чтения данных из канала
	numWorkers := 5 // Количество воркеров
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerNum int) {
			defer wg.Done()
			for {
				select {
				case data := <-dataChan:
					// Выводим данные в stdout
					fmt.Printf("Воркер %d: %s\n", workerNum, data)
				case <-doneChan:
					return
				}
			}
		}(i)
	}

	// Ожидаем сигналы об окончании работы программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнала об окончании работы программы или воркеров
	select {
	case <-sigChan:
		// Получен сигнал об окончании работы программы
		fmt.Println("Выход из программы")
	case <-doneChan:
		// Получен сигнал об окончании работы воркеров
		fmt.Println("Воркер завершил работу")
	}

	// Останавливаем работу воркеров
	close(doneChan)

	// Ожидаем завершения работы воркеров
	wg.Wait()
	fmt.Println("Программы завершила работу")
}
