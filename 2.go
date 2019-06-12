package main
/*AIM-Write  program that prints all the prime numbers between ‘x’ and ‘y’, inclusive of ‘x’ and ‘y’ without using a loop where ‘x’ and ‘y’ will be specified by the user.*/

import (
	"fmt"
	"strings"
)



func parse (s string ) int {
	sum := 0
	for i :=0; i<len(s); i++{
		if s[i] >= 48 && s[i] <= 57{
			sum += int(s[i]) - 48
		}
	}

	//if(sum>26){
	//	sum = sum % 26
	//}

	return sum
}

func main() {
	var x string
	x = "Parul9 raich"
	sarr := strings.Fields(x)
	for i := 0; i < len(sarr); i++{
		if parse(sarr[i]) + 96 != 96 {
			fmt.Print(string(parse(sarr[i]) + 96))
		} else if  parse(sarr[i]) + 96 == 96{
			fmt.Println("No code encoded")
		}

		//fmt.Println()
	}

}
