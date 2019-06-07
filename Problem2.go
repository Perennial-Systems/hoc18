package main

import (
	"log"
	"strconv"
	"unicode"
)

func printChar(i int)  {
	println(rune('A' - 1 + i))
}

func IsDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main(){
	var code string
	code = "aWk60#ase2> o_0^4alis5L"
	//fmt.Println(code[22:23])
	for i := 0; i < len(code); i++ {
			for  j := 1; j <= len(code); j++ {
				var sum []int
				if IsDigit(code[i:j]) == true{
					var digit int
					digit, err := strconv.Atoi(code[i:j])
					if err != nil  {
						log.Fatal(0)
					}
					sum = append(sum, digit)
				} else if code[i:j] == " " {
					var x int
					for x, v := range sum{
						x += v
					}
					sum = nil
					printChar(x)
			}
		}
	}
}
