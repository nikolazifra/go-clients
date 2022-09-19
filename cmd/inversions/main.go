package main

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
		for i := n; i >= 0; i-- {
			result = append(result, i)
		}
		r <- result

	}(r, 2000000)

	a := <-r
	//start := time.Now()
	MergeSort(a)
	//elapsed := time.Since(start)
	//fmt.Println(i5, elapsed)
}
