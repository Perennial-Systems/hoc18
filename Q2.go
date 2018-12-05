package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	line := "aWk60#ase2> o_2^293alis874L"
	words :=strings.Split(line," ")
	var resStr string
	for _,word:=range words{
		char:=getTotal(word)
		resStr=resStr+string(char)
	}
	fmt.Printf("%s",resStr)
}

func getTotal(word string)(res int){
	x:=0
	var total int
	for x < len(word){
		ch := string([]rune(word)[x])
		i, err := strconv.Atoi(ch)
		if err == nil {
			total = total+i
		}
		x+=1
	}
	if total==0{
		return 32
	}else{
		if total>26{
			res = (total%26)+64
			return res
		}else {
			return total+64
		}
	}

}