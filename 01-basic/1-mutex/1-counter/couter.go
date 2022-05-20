package main

import (
	"fmt"
	"sync"
)

const routineNum = 10
const incNum = 100000

func main() {
	//互斥锁保护计数器
	var mutex sync.Mutex
	var count = 0
	//使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(routineNum)

	for i := 0; i < routineNum; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < incNum; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	//主协程等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}