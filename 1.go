package main
/*AIM-Write  program that prints all the prime numbers between ‘x’ and ‘y’, inclusive of ‘x’ and ‘y’ without using a loop where ‘x’ and ‘y’ will be specified by the user.*/

import "fmt"

func primenumber(y int64,i int64) int64{
 if i == int64(1) {
 return 1} else if i > 1{
 if y%i == 0 {
 return 0} else {
return primenumber(y, i-1)
}
}
return -1
}

func findprime(n1 int64,n2 int64){
	var x, i int64
	x = n1
	i = x / 2

	goto CHECK
CHECK:
	if x <= n2 {
		isPrime := primenumber(x, i)
		if isPrime == 1 {
			fmt.Println(x)
			x++
			goto CHECK
		} else {
			x++
			goto CHECK
		}
	}

}





func main() {
	var n1, n2 int64
	fmt.Println("Enter the first number")
	fmt.Scan(&n1)
	fmt.Println("enter the second number")
	fmt.Scan(&n2)

	fmt.Println("numbers are:")
	findprime(n1,n2)}
