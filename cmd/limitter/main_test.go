package main

import (
	"fmt"
	"sync"
	"testing"
)

func Test_RaceCondition1(t *testing.T) {
	rateLimiter := make(chan bool)
	sm := make(map[string]string, 8000)
	var mutex *sync.Mutex = &sync.Mutex{}
	for i := 0; i < 8000; i++ {
		go sendTick(rateLimiter, i, sm, mutex)
	}
}

func Test_RaceCondition2(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i1 int) {
			defer wg.Done()
			fmt.Println(i1)
		}(i)

	}

	wg.Wait()
	fmt.Println("Done")

}
