package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("Enter Two Strings seprating a space ")
	var s1, s2 string
	fmt.Scanf("%s %s", &s1, &s2)

	re := regexp.MustCompile("[0-9]+")
	str1 := make([]string, 100)
	str1 = re.FindAllString(s1, -1)

	str2 := make([]string, 100)
	str2 = re.FindAllString(s2, -1)

	//arr := make([]string, 100)
	//w1 := arr[0]
	//w2 := arr[]
	//len(w1) w1[1]

	var i, num int

	for i, num := range str1 {
		var sum1 int
		sum1 = 0
		sum1 = sum1 + []int(str1[i])
	}
}
