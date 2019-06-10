package main

import "fmt"

func main() {
	var i, f int
	fmt.Println("Enter the Range")
	_, _ = fmt.Scan(&i, &f)
	print(2, 6)
}

func print(i int, f int) {

check1:
	if i <= f {
		var num, fact, k int
		num = i
		fact = 0
		k = 2
	check:
		if k <= (num / 2) {
			if num%k == 0 {
				fact++
			}
			k++
			goto check
		}
		if fact == 0 {
			fmt.Println(i)
		}

		i++
		goto check1
	}
}
