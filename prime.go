package main
import "fmt"

func isPrime(min int, max int) {

	var i = 2
if min <= 2 {
	fmt.Println("%d",min);
}
	Loop:
if min % i != 0 {
	i++
	if i<min/2 {
	goto Loop
	}
	fmt.Println("%d",min);				//Print number if no factor found

}
	var m = min+1
if m<=max {
	isPrime(m, max)
}
}

func main() {
var min,max int

fmt.Printf("Enter the numbers(min and max):")
fmt.Scanf("%d %d",&min,&max)
isPrime(min,max)
}
