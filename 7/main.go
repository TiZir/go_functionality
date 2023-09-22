package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	m := make(map[int]int)
	mu := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			m[val] = val
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println(len(m))
}
