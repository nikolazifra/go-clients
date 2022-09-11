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

	"github.com/benchkram/errz"
	err "github.com/pkg/errors"
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
		errz.Fatalf(err, "File %s could not be found.", fileName)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineRead <- scanner.Text()
			fmt.Println(">", scanner.Text())
		}
		file.Close()
		cancel()
	}()

	// goroutine to read file line by line and passing to channel to print
	wg.Add(1)
	go func(context.Context) {
		defer wg.Done()
		var b strings.Builder
		for {
			select {
			case <-ctx.Done():
				errz.Log(err.Wrap(ctx.Err(), "Cancel encountered"))
				fmt.Println(b.String())
				return
			case line := <-lineRead:
				var l strings.Builder
				l.WriteString(line)
				l.WriteString("\n")
				fmt.Println("*", l.String())
				b.WriteString(l.String())
			case <-time.After(time.Microsecond):
				fmt.Println("Chilling out for 1 microsec...")
			}
		}
	}(ctx)

	wg.Wait()
	test := make(chan int, 3)
	test <- 2
	test <- 4
	test <- 5
	fmt.Println("Terminating...")
	// type function tests..
	type tFunc func(int, int) int
	var f tFunc = func(x int, y int) int { return x + y }
	fmt.Println(f(1, 0))

}
