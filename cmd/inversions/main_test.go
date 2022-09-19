package main

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_Merge2(t *testing.T) {
	a := []int{3, 4}
	b := []int{1, 2}
	want := []int{1, 2, 3, 4}
	if got, x := merge(a, b, 2); !reflect.DeepEqual(want, got) || x != 4 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 4)
	}

}

func Test_Merge1(t *testing.T) {
	a := []int{3}
	b := []int{1}
	want := []int{1, 3}
	if got, x := merge(a, b, 1); !reflect.DeepEqual(want, got) || x != 1 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 1)
	}

}

func Test_Merge3(t *testing.T) {
	a := []int{2, 3, 4}
	b := []int{1}
	want := []int{1, 2, 3, 4}
	if got, x := merge(a, b, 3); !reflect.DeepEqual(want, got) || x != 3 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 3)
	}

}

func Test_Merge4(t *testing.T) {
	a := []int{1}
	b := []int{3}
	want := []int{1, 3}
	if got, x := merge(a, b, 1); !reflect.DeepEqual(want, got) || x != 0 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 0)
	}

}

func Test_Merge5(t *testing.T) {
	a := []int{1}
	b := []int{2, 3, 4}
	want := []int{1, 2, 3, 4}
	if got, x := merge(a, b, 1); !reflect.DeepEqual(want, got) || x != 0 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 0)
	}

}

func Test_Merge6(t *testing.T) {
	a := []int{3}
	b := []int{2, 3, 4}
	want := []int{2, 3, 3, 4}
	if got, x := merge(a, b, 1); !reflect.DeepEqual(want, got) || x != 1 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 1)
	}

}

func Test_Merge7(t *testing.T) {
	a := []int{2}
	b := []int{3}
	want := []int{2, 3}
	if got, x := merge(a, b, 1); !reflect.DeepEqual(want, got) || x != 0 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 0)
	}

}

func Test_MergeSort1(t *testing.T) {
	a := []int{4, 3, 2, 1}
	want := []int{1, 2, 3, 4}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 6 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 6)
	}
}

func Test_MergeSort2(t *testing.T) {
	a := []int{2, 3, 1}
	want := []int{1, 2, 3}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 2 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 2)
	}
}

func Test_MergeSort3(t *testing.T) {
	a := []int{8, 4, 2, 1}
	want := []int{1, 2, 4, 8}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 6 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 6)
	}
}

func Test_MergeSort4(t *testing.T) {
	a := []int{8, 4, 2, 1}
	want := []int{1, 2, 4, 8}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 6 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 6)
	}
}

func Test_MergeSort5(t *testing.T) {
	a := []int{3, 1, 2}
	want := []int{1, 2, 3}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 2 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 2)
	}
}

func Test_MergeSort6(t *testing.T) {
	a := []int{3, 1, 2}
	want := []int{1, 2, 3}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 2 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 2)
	}
}

func Test_MergeSort7(t *testing.T) {
	a := []int{1, 20, 6, 4, 5}
	want := []int{1, 4, 5, 6, 20}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 5 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 5)
	}
}

func Test_MergeSort8(t *testing.T) {
	a := []int{-1, 6, 3, 4, 7, 4}
	want := []int{-1, 3, 4, 4, 6, 7}
	if got, x := MergeSort(a); !reflect.DeepEqual(want, got) || x != 4 {
		t.Errorf("Merge() = want %v, got %v, inversions=%d, but expected count was %d", want, got, x, 4)
	}
}

func BenchmarkMergeSort_10000(b *testing.B) {
	r := make(chan []int)
	const n int = 10000
	for i := 0; i < b.N; i++ {

		go func(r chan []int, n int) {
			result := make([]int, 0, n-1)
			for j := n; j >= 0; j-- {
				result = append(result, j)
			}
			r <- result
		}(r, n)
		a := <-r
		start := time.Now()
		_, i5 := MergeSort(a)
		elapsed := time.Since(start)
		fmt.Println(i5, elapsed)
	}
}

func BenchmarkMergeSort_100000(b *testing.B) {
	r := make(chan []int)
	const n int = 100000
	for i := 0; i < b.N; i++ {

		go func(r chan []int, n int) {
			result := make([]int, 0, n-1)
			for j := n; j >= 0; j-- {
				result = append(result, j)
			}
			r <- result
		}(r, n)
		a := <-r
		start := time.Now()
		_, i5 := MergeSort(a)
		elapsed := time.Since(start)
		fmt.Println(i5, elapsed)
	}
}

func BenchmarkMergeSort_1000000(b *testing.B) {
	r := make(chan []int)
	const n int = 1000000
	for i := 0; i < b.N; i++ {

		go func(r chan []int, n int) {
			result := make([]int, 0, n-1)
			for j := n; j >= 0; j-- {
				result = append(result, j)
			}
			r <- result
		}(r, n)
		a := <-r
		start := time.Now()
		_, i5 := MergeSort(a)
		elapsed := time.Since(start)
		fmt.Println(i5, elapsed)
	}
}
