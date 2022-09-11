package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	X int    = 1
	Y string = ""
)

func main() {

	fmt.Printf("Cores:%d\n", runtime.NumCPU())
	ctx, cancel := context.WithCancel(context.Background())
	lineRead := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var fileName = "sample-file.txt"
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("failed opening file: %s", err)
			return
		} else {
			fmt.Println("Ok!!!")
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineRead <- scanner.Text()
			fmt.Println(">", scanner.Text())
			//time.Sleep(200 * time.Millisecond)
		}
		file.Close()
		cancel()
	}()

	// goroutine to read file line by line and passing to channel to print
	wg.Add(1)
	go func() {
		defer wg.Done()
		var b strings.Builder
		for {
			select {
			case <-ctx.Done():
				fmt.Println("process stopped. reason: ", ctx.Err())
				fmt.Println(b.String())
				return
			case line := <-lineRead:
				fmt.Println("*", line)
				b.WriteString(line)
				b.WriteString("\n")
			case <-time.After(time.Microsecond):
				fmt.Println("Chilling out for 1 microsec...")
			}
		}
	}()

	wg.Wait()
	test := make(chan int, 3)
	test <- 2
	test <- 4
	test <- 5
	fmt.Println("Terminating...")

	type tFunc func(int, int) int
	var f tFunc = func(x int, y int) int { return x + y }
	fmt.Println(f(1, 0))

}
