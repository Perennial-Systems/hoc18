package main

import "fmt"

func main() {
	var i, f int
	fmt.Println("Enter the Range")
	fmt.Scanf("%d%d", &i, &f)
	print(i, f)
}

func print(i int, f int) {

	if i <= f {
		var num, fact, k int
		num = i
		fact = 0
		k = 2
		if k <= (num / 2) {
			if num%k == 0 {
				fact++
				k++
			}
			i++
			if fact >= 1 {
				fmt.Println(i)
			}
		}
	}
}
