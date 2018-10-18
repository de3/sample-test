package main

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n%2 == 0 {
		return false
	}
	if n%3 == 0 {
		return false
	}

	return true
}

func sumSlice(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}

	return sum
}

func nextN(n int) int {
	return n/2 + n%2
}

func c(n int) {
	i := n
	tmp := make([]int, 0)
	final := make([][]int, 0)
	for i > 1 {
		if sumSlice(tmp) == n {
			final = append(final, tmp)
			tmp = make([]int, 0)
		}

		if isPrime(i) && i == n {
			tmp = append(tmp, i)
			i = nextN(i)
			continue
		}

		if isPrime(i) {
			tmp = append(tmp, i)
		}
		i = nextN(i)
	}

	if len(tmp) > 0 && sumSlice(tmp) == n {
		final = append(final, tmp)
	}

	fmt.Println(final)
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	c(n)
}
