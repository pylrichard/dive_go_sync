package main

import (
	"fmt"
	"sync"
)

const routineNum = 10
const incNum = 100000

type Counter struct {
	Type int
	Name string

	mutex sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mutex.Lock()
	c.count++
	c.mutex.Unlock()
}

func (c *Counter) Count() uint64 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return c.count
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(routineNum)

	for i := 0; i < routineNum; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < incNum; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}