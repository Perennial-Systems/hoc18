package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)
	fmt.Print("Prime numbers are: ")

	for i := num1; i <= num2; i++ {
		primeNo(i, i/2)
	}
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
