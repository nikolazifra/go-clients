package main

import (
	"fmt"
)

func primeNums(start, end int) []int {
	var result []int = make([]int, 0, end-start)
	for j := start; j <= end; j++ {
		c := 0
		for i := 1; i <= j; i++ {
			if j%i == 0 {
				c++
			}
			if c >= 3 {
				break
			}
		}
		if c == 2 {
			result = append(result, j)
		}
	}
	return result
}

func d(x int, primes []int) (int, int, int) {
	for i, p := range primes {
		if x == p {
			fmt.Println(1, x)
			return p, p, x
		} else if p > x && i > 0 && i <= len(primes)-1 {
			fmt.Println(2, x)
			return primes[i-1], primes[i], x
		} else if x > p && i == len(primes)-1 {
			// if 99 or 98 is supplied
			fmt.Println(3, x)
			return p, p, x
		}
	}
	return 0, 0, 0
}

func calc(l, r, a int) int {
	return (l + r) * a
}

func SpecialArray(n int, arr []int) []int {
	var result []int = make([]int, 0, n)
	primes := primeNums(1, 99)
	fmt.Println(primes)
	for _, x := range arr {
		result = append(result, calc(d(x, primes)))
	}
	return result
	//Insert your code here
}

func main() {
	fmt.Println("Result is", SpecialArray(3, []int{2, 10, 23}))
	fmt.Println("Result is", SpecialArray(5, []int{4, 5, 10, 14, 8}))
	fmt.Println("Result is", SpecialArray(3, []int{3, 11, 97, 99}))
	fmt.Println("Result is", SpecialArray(3, []int{12, 3, 43, 91, 72, 6, 53, 21, 9}))
}
