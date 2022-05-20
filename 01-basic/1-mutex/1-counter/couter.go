package main

import (
	"fmt"
	"sync"
)

const routineNum = 10
const incNum = 100000

func main() {
	var count = 0
	//使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(routineNum)

	for i := 0; i < routineNum; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < incNum; j++ {
				count++
			}
		}()
	}
	//主协程等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}