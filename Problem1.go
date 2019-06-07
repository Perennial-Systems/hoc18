package main

import (
	"fmt"
	"os"
)

func primeNo(num int, i int) int {
	if i == 1 {
		return 1
	} else if i > 1{
		if num%i == 0 {
			return 0
		} else {
			return primeNo(num, i-1)
		}
	}
	return -1
}

func main(){
	var x, y int
	fmt.Println("Enter x:")
	fmt.Scanln(&x)
	fmt.Println("Enter y:")
	fmt.Scanln(&y)
	if x >= y{
		fmt.Println("Please enter value of x greater than y ")
		os.Exit(0)
	}
	if x <= 1 || y <= 1{
		fmt.Println("Please enter positive integer values of x and y greater than 1")
		os.Exit(0)
	}

	fmt.Println("Prime numbers between", x,"&",  y, "are" )
	n := x
	i := n/2
	goto LABEL
	LABEL: if n <= y{
		isPrime := primeNo(n, i)
		if isPrime == 1 {
			fmt.Println(n)
			n++
			goto LABEL
		} else{
			n++
			goto LABEL
		}
	}
}
