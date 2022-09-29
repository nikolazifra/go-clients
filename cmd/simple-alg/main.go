package main

import (
	"bytes"
	"fmt"
	"strings"
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

func toMap(a []int) map[int]int {
	result := make(map[int]int, len(a))
	for _, val := range a {
		if cnt, found := result[val]; found {
			result[val] = cnt + 1
		} else {
			result[val] = 1
		}

	}
	return result

}

func FirstUnique(A []int) int {
	m := toMap(A)
	for _, val := range A {
		if cnt, found := m[val]; found && cnt == 1 {
			return val
		}
	}
	return -1
}

func Solution1(A []int) int {
	return FirstUnique(A)
}

func Solution2(A, B int) string {
	// if A>B append "aab"
	// if B>A append "bba"
	// if B=A append "ab"
	// aab (3,2) // aab (1,1) //ab
	var sb bytes.Buffer
	for 0 < A || 0 < B {
		if A > B {
			if A > 0 {
				sb.WriteString("a")
				A -= 1
			}
			if A > 0 {
				sb.WriteString("a")
				A -= 1
			}
			if B > 0 {
				sb.WriteString("b")
				B -= 1
			}
		} else if B > A {
			if B > 0 {
				sb.WriteString("b")
				B -= 1
			}
			if B > 0 {
				sb.WriteString("b")
				B -= 1
			}
			if A > 0 {
				sb.WriteString("a")
				A -= 1
			}
		} else {
			if A > 0 {
				sb.WriteString("a")
				A -= 1
			}
			if B > 0 {
				sb.WriteString("b")
				B -= 1
			}
		}
	}
	return sb.String()
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func ReverseWords(s string) string {
	r := strings.Split(s, " ")
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return strings.Join(r, " ")
}

func main() {
	fmt.Println("Result is", SpecialArray(3, []int{2, 10, 23}))
	fmt.Println("Result is", SpecialArray(5, []int{4, 5, 10, 14, 8}))
	fmt.Println("Result is", SpecialArray(3, []int{3, 11, 97, 99}))
	fmt.Println("Result is", SpecialArray(3, []int{12, 3, 43, 91, 72, 6, 53, 21, 9}))
	fmt.Println("First unique ", FirstUnique([]int{1, 2, 1, 4, 1000000000, 4, 2}))
	fmt.Println("String generation ", Solution2(1, 3))
	fmt.Println("Reverse string", Reverse("nikola"))
	fmt.Println("Reverse string", ReverseWords("quick brown fox has jumped"))
}
