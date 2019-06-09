package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)
	fmt.Print("Prime numbers are: ")

	rng(num1, num2)
}

func primeNo(i, j int) {
	if j == 1 {
		fmt.Print(i, " ")
	} else {
		if i%j == 0 {
			fmt.Print("")
		} else {
			primeNo(i, j-1)
		}
	}
}

func rng(num1, num2 int) {
	if num1 <= num2 {
		primeNo(num1, num1/2)
		rng(num1+1, num2)
	}
}
