// select.go
package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	REQUESTS_PS int = 100
)

func doSomethingHighPriority() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("very important stuff done at", time.Now())
}

func doSomethingLowPriority() {
	time.Sleep(time.Millisecond)
	fmt.Println("default action")
}

func sendTick(rateLimiter chan<- bool, k int, m map[string]string, mutex *sync.Mutex) {
	b := strings.Builder{}
	ks := strconv.Itoa(k)
	b.WriteString("hi_")
	b.WriteString(ks)
	mutex.Lock()
	fmt.Println("-- Lock acquired by goroutine:", k)
	m[ks] = b.String()
	fmt.Println("-- Releasing mutex lock")
	mutex.Unlock()

	rate := time.Tick(time.Second / 10)
	for range rate {
		rateLimiter <- true
		//fmt.Println("Limmiting rate at", time.Now())
	}
}

func receive(rateLimiter <-chan bool) {

	fmt.Println("Passed requests limit check...")
	select {
	case <-rateLimiter:
		doSomethingHighPriority()
	default:
		doSomethingLowPriority()
	}
}

func main() {
	rateLimiter := make(chan bool, 3)
	sm := make(map[string]string, 4000000)
	var mutex *sync.Mutex = &sync.Mutex{}
	for i := 0; i < 2; i++ {
		go sendTick(rateLimiter, i, sm, mutex)
	}
	limitter := rate.NewLimiter(4, 5)
	for {
		//fmt.Println("Waiting to be processed:", len(rateLimiter))
		if !limitter.Allow() {
			fmt.Println("Dropping requests...")
			//time.Sleep(time.Second / 10)
			time.Sleep(time.Millisecond * 100)
			continue
		} else {
			receive(rateLimiter)
		}

		// Sleep to not perform the default action too much
	}
}
