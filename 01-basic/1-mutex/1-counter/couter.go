package main

import (
	"fmt"
	"sync"
)

const routineNum = 10
const incNum = 100000

type Counter struct {
	sync.Mutex
	Count uint64
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(routineNum)

	for i := 0; i < routineNum; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < incNum; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}