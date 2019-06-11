package main

import (
	"fmt"
	"strings"
)

func main() {

	typed := "aWk999#ase> o_0^4alis5L" //Hardcoded can be removed by removing following commented statements
	//fmt.Println("Enter the encoded string:x")
	//fmt.Scan(&typed)

	brokenTyped := strings.Fields(typed)

	for i := 0; i < len(brokenTyped); i++ {
		j := decode(brokenTyped[i])

		j = j%26 // this is to loop back if value > 26

		if j == 0 {
			j = 32
		} else {
			j += 96 //to make 1 == ascii(97) is 'a'
		}
		fmt.Print(string(j))
	}
}

//ascii 48 to 57 is 0 to 9
//32 ascii is space
//97 is a

func decode(x string) int {
	j := 0
	for i := 0; i < len(x); i++ {
		if x[i] >= 48 && x[i] <= 57 {
			j += int(x[i]) - 48
		}
	}
	return j
}
