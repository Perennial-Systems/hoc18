package main

import "fmt"

func primeNo(y int64,i int64) int64{
		if i == int64(1) {
		return 1} else if i > 1{
		if y%i == 0 {
		return 0} else {
		return primeNo(y, i-1)
		}
	}
		return -1
}

func main(){
	var n1, n2, retval int64
	fmt.Println("Enter the first number")
	fmt.Scan(&n1)
	fmt.Println("enter the second number")
	fmt.Scan(&n2)

	if n1<n2 {
		x := n1
		CHECK: if x>=n1 && x<= n2{
			retval = primeNo(x,x/2)
			if retval == 1{
				println(x)
			}
			x++
			goto CHECK
		}
	}

	
	if  n1 > n2{
		n1,n2 = n2,n1
		x := n1
		CHECKIF: if x>=n1 && x<= n2{
			retval = primeNo(x,x/2)
			if retval == 1{
				println(x)
			}
			x++
			goto CHECKIF
		}
		
	}



}
