package main

import (
	"fmt"
	"time"
)

func merge(a []int, b []int) ([]int, int) {
	result := make([]int, 0, len(a)+len(b))
	//result := []int{}
	i := 0
	j := 0
	inv := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
			inv += len(a) - i
		}
	}
	for ; i < len(a); i++ {
		result = append(result, a[i])
	}
	for ; j < len(b); j++ {
		result = append(result, b[j])
	}
	return result, inv
}

func MergeSortP(items []int) chan []int {
	x := len(items) / 2
	l_chan := make(chan []int, 2)
	r_chan := make(chan []int, 2)
	r_res := make(chan []int, 2)
	go func(a chan<- []int, c int) {
		l_arr, l := MergeSort(items[:c])
		a <- l_arr
		a <- []int{l}
	}(l_chan, x)
	go func(a chan<- []int, c int) {
		r_arr, r := MergeSort(items[c:])
		a <- r_arr
		a <- []int{r}
	}(r_chan, x)
	result, inv := merge(<-l_chan, <-r_chan)
	l := (<-l_chan)[0]
	r := (<-r_chan)[0]
	if l+r+inv > 100000000 {
		r_res <- result
		r_res <- []int{-1}
		return r_res
	}
	r_res <- result
	r_res <- []int{l + r + inv}
	return r_res
}

func MergeSort(items []int) ([]int, int) {
	if len(items) < 2 {
		return items, 0
	}
	x := len(items) / 2
	l_arr, l := MergeSort(items[:x])
	r_arr, r := MergeSort(items[x:])
	result, m := merge(l_arr, r_arr)
	if l+r+m > 100000000 {
		return result, -1
	}
	return result, l + r + m
}

func main() {

	r := make(chan []int)
	go func(r chan []int, n int) {
		result := make([]int, 0, n-1)
		for i := n; i >= 1; i-- {
			result = append(result, i)
		}
		r <- result

	}(r, 100000000)

	a := <-r
	t := time.Now()
	c := MergeSortP(a)
	x := <-c
	y := <-c
	elapsed := time.Since(t)
	fmt.Println("Parallel", elapsed, len(x), y[0])
	t = time.Now()
	a, inv := MergeSort(a)
	elapsed = time.Since(t)
	fmt.Println("Sequential", elapsed, len(a), inv)
}
