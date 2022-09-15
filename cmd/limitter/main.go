// select.go
package main

import (
	"fmt"
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

func sendTick(rateLimiter chan<- bool) {
	rate := time.Tick(time.Second)
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
	rateLimiter := make(chan bool)
	for i := 0; i < 1000; i++ {
		go sendTick(rateLimiter)
	}
	limitter := rate.NewLimiter(10, 1)
	for {
		//fmt.Println("Waiting to be processed:", len(rateLimiter))
		if !limitter.Allow() {
			fmt.Println("Dropping requests...")
			time.Sleep(time.Second)
			continue
		}
		receive(rateLimiter)
		// Sleep to not perform the default action too much
	}
}
