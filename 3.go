package main

import "fmt"

func findZprimes(k int , start int, end int){
	var primeFact []int
	var capacity int
	for i:=start ; i<= end;i++ {
		primeFact,capacity = primeFactor(i)
		if k == capacity{
			fmt.Printf("the number %v is a z  prime number with z value = %v with %v factors",i,k,primeFact)
			fmt.Println()
		}
	}
}

func primeFactor(y int) (pf []int,cap int){
	//n := y
	//cap = 0
	for y%2 == 0 {
		pf = append(pf, 2)
		cap++
		y = y / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= y; i = i + 2 {
		// while i divides n, append i and divide n
		for y%i == 0 {
			pf = append(pf, i)
			cap++
			y = y / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if y > 2 {
		pf = append(pf, y)
		cap++
	}

	return
}

func main() {
	var k,start,end int
    fmt.Println("input the Z value")
	fmt.Scan(&k)
	fmt.Println("input the start value")
	fmt.Scan(&start)
	fmt.Println("input the end value")
	fmt.Scan(&end)
	findZprimes(k,start,end)
}