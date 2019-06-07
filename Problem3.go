package main

import (
	"fmt"
)

func main() {
	var k, start, end int
	fmt.Println("Enter k, start, end:")
	fmt.Scan(&k, &start, &end)
	fmt.Println(findZprimes(k, start, end))
}

func findZprimes(k, start, end int) (ans []int) {
	for i := start; i <= end; i++ {
		j := retNum(i)
		if j == k {
			ans = append(ans, i)
		}
	}
	return ans
}

func retNum(i int) int {
	n := 0
	for i/2 == 0 {
		i = i / 2
		n++
	}

	return n
}
