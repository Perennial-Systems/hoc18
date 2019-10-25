package main

import (
	"fmt"
)

func main() {
	var k, start, end int
	fmt.Println("Enter k, start, end:")
	fmt.Scan(&k, &start, &end)
	fmt.Print("Numbers in range which are ",k,"-prime are:")
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
	for j := 2; j <= i/2; j++{
		for i%j == 0 {
			i = i / j
			n++
		}
	}
	if i > 2{
		n++
	}

	return n
}
