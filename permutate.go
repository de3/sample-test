package main

import (
	"fmt"
	"os"
	"strconv"
)

func makeSlice(n int) []int {
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, i+1)
	}

	return res
}

func permutate(s []int) [][]int {
	result := make([][]int, 0)

	var p func([]int, int)
	p = func(a []int, k int) {

		if k == len(a) {
			result = append(result, append([]int{}, a...))
			return
		}

		for i := k; i < len(s); i++ {
			a[i], a[k] = a[k], s[i]
			p(a, k+1)
			a[i], a[k] = a[k], s[i]
		}

	}
	p(s, 0)

	return result
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])

	s := makeSlice(n)
	result1 := permutate(s)
	fmt.Println(result1)
}
