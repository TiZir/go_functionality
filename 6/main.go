package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Гоурутина осанавливается по сигналу остановки через канал
func worker1(stopCh chan bool) {
	for {
		select {
		case <-stopCh:
			return
		default:
			fmt.Println("Воркер 1...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Гоурутина осанавливается по закрытию контекста
func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("Воркер 2...")
			time.Sleep(1 * time.Second)
		}
	}
}

// Гоурутина осанавливается по концу выполениня с помощью sync.WaitGroup
func worker3(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("Воркер 3...")
		time.Sleep(1 * time.Second)
	}
}

// Гоурутина осанавливается по концу выполениня с помощью таймера
func worker4(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Воркер 4...")
	for {
		select {
		case <-time.After(time.Second * 5):
			fmt.Println("Остановка воркера 4...")
			return
		}
	}
}

func main() {

	//Воркер 1
	stopCh := make(chan bool)   // канал для остановки
	go worker1(stopCh)          // запуск воркера
	time.Sleep(5 * time.Second) // ждем 5 секунд(итерация)
	fmt.Println("Остановка воркера 1...")
	stopCh <- true // остановка воркера
	time.Sleep(2 * time.Second)

	//Воркер 2
	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("Остановка воркера  2...")
	cancel() // закрываем контекст
	time.Sleep(2 * time.Second)

	//Воркер 3
	var wg sync.WaitGroup
	wg.Add(1)
	go worker3(&wg)
	time.Sleep(5 * time.Second)
	fmt.Println("Остановка воркера 3...")

	//Воркер 4
	wg.Add(1)
	go worker4(&wg)

	wg.Wait()
	fmt.Println("Main function stopped")
}
