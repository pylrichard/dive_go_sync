package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	mutexLocked = 1 << iota
	mutexWakened
	mutexStarving
	mutexWaiterShift = iota
)

type Mutex struct {
	sync.Mutex
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked)
}

func (m *Mutex) Count() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	waiter := v >> mutexWaiterShift
	waiter = waiter + (v & mutexLocked)

	return int(waiter)
}

func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))

	return state & mutexLocked == mutexLocked
}

func (m *Mutex) IsWakened() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))

	return state & mutexWakened == mutexWakened
}

func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))

	return state & mutexStarving == mutexStarving
}

func main() {
	try()

	count()
}

func try() {
	var mutex Mutex
	go func() {
		mutex.Lock()
		time.Sleep(time.Second)
		mutex.Unlock()
	}()

	time.Sleep(time.Second)

	ok := mutex.TryLock()
	if ok {
		fmt.Println("get the lock")
		mutex.Unlock()

		return
	}

	fmt.Println("can't acquire the lock'")
}

func count() {
	var mutex Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock()
			time.Sleep(time.Second)
			mutex.Unlock()
		}()
	}

	time.Sleep(time.Second)

	fmt.Printf("waiting: %d, wakened: %t, starving: %t\n", mutex.Count(), mutex.IsWakened(), mutex.IsStarving())
}