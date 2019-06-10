package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	fmt.Println("Enter Two Strings seprating a space ")
	var s1, s2 string
	_, _ = fmt.Scanf("%s %s", &s1, &s2)

	re := regexp.MustCompile("[0-9]+")
	str1 := make([]string, 100)
	str1 = re.FindAllString(s1, -1)
	sum1 := 0
	for i := 0; i < len(str1); i++ {
		k, _ := strconv.Atoi(str1[i])
		sum1 += k
		sum1 = sum1 % 26
	}
	//fmt.Println(sum1)
	if sum1 == 0 {
		fmt.Print("No number in a string")
	} else {
		fmt.Print(string(sum1 + 96))
	}
	str2 := make([]string, 100)
	str2 = re.FindAllString(s2, -1)
	sum2 := 0
	for j := 0; j < len(str2); j++ {
		k, _ := strconv.Atoi(str2[j])
		sum2 += k
		sum2 = sum2 % 26
	}
	//fmt.Println(sum2)
	if sum2 == 0 {
		fmt.Print(string(sum2 + 96))
	} else {
		fmt.Println("No number in a string")
	}

}
