package main

import "fmt"

func merge(a []int, b []int, m int) ([]int, int) {
	result := make([]int, 0, len(a)+len(b))
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
			inv += m - i
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

func MergeSort(items []int) ([]int, int) {
	if len(items) < 2 {
		return items, 0
	}
	x := len(items) / 2
	first, l := MergeSort(items[:x])
	second, r := MergeSort(items[x:])
	result, m := merge(first, second, x)
	return result, l + r + m
}

func main() {

	r := make(chan []int)
	go func(r chan []int, n int) {
		result := make([]int, 0, n-1)
		for i := n; i >= 0; i-- {
			result = append(result, i)
		}
		r <- result

	}(r, 1000)

	s1, i1 := MergeSort([]int{-1, 6, 3, 4, 7, 4})
	fmt.Println(s1, i1)

	s2, i2 := MergeSort([]int{1, 20, 6, 4, 5})
	fmt.Println(s2, i2)

	s3, i3 := MergeSort([]int{3, 1, 2})
	fmt.Println(s3, i3)

	s4, i4 := MergeSort([]int{8, 4, 2, 1})
	fmt.Println(s4, i4)
	a := <-r
	s5, i5 := MergeSort(a)
	fmt.Println(s5, i5)
}
