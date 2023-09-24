package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f          Function
	InProgress map[int]bool
	IsPending  map[int][]chan int
	cache      map[int]FunctionResult
	lock       sync.RWMutex
}

type Function func(key int) (int, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:          f,
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
		cache:      make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()
	if !exists {
		result.value, result.err = m.Work(key)
		m.lock.Lock()
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}

func (m *Memory) Work(job int) (int, error) {
	m.lock.RLock()
	exists := m.InProgress[job]
	if exists {
		m.lock.RUnlock()
		response := make(chan int)
		defer close(response)

		m.lock.Lock()
		m.IsPending[job] = append(m.IsPending[job], response)
		m.lock.Unlock()
		resp := <-response
		return resp, nil
	}
	m.lock.RUnlock()

	m.lock.Lock()
	m.InProgress[job] = true
	m.lock.Unlock()

	result, _ := m.f(job)

	m.lock.RLock()
	pendingWorkers, exists := m.IsPending[job]
	m.lock.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
	}
	m.lock.Lock()
	m.InProgress[job] = false
	m.IsPending[job] = make([]chan int, 0)
	m.lock.Unlock()
	return result, nil
}

func GetFibonacci(n int) (int, error) {
	fmt.Println("Se ejecuto GetFibonacci")
	return Fibonacci(n), nil
}

func main() {
	timeInit := time.Now()
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38, 40, 40, 41, 43, 43, 43}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)
	}
	wg.Wait()
	fmt.Printf("tiempo final %s\n", time.Since(timeInit))
}
