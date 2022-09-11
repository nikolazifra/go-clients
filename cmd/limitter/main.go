// select.go
package main

import (
	"fmt"
	"time"
)

const (
	REQUESTS_PS int = 100
)

func doSomethingHighPriority() {
	fmt.Println("very important stuff done at", time.Now())
}

func doSomethingLowPriority() {
	fmt.Println("default action")
}

func sendTick(rateLimiter chan<- bool) {
	rate := time.Tick(time.Second / time.Duration(REQUESTS_PS))
	for range rate {
		rateLimiter <- true
	}
}

func receive(rateLimiter <-chan bool) {
	select {
	case <-rateLimiter:
		doSomethingHighPriority()
	default:
		doSomethingLowPriority()
	}
}

func main() {
	rateLimiter := make(chan bool, REQUESTS_PS)
	go sendTick(rateLimiter)
	for {
		receive(rateLimiter)
		time.Sleep(2 * time.Millisecond) // Sleep to not perform the default action too much
	}
}
