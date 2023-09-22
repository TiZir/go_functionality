package main

import (
	"fmt"
	"time"
)

func sleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C // ждем пока не выйдет время
	// функция завершает работу программа идет дальше
}

func main() {
	sleep(5 * time.Second)
	fmt.Println("Конец")
}
