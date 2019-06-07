package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {

	fmt.Println("Enter Two Strings seprating a space ")
	var s1, s2 string
	fmt.Scanf("%s %s", &s1, &s2)

	re := regexp.MustCompile("[0-9]+")
	str1 := make([]string, 100)
	str1 = re.FindAllString(s1, -1)
	sum1 := 0
	for i := 0; i < len(str1); i++ {
		k, _ := strconv.Atoi(str1[i])
		sum1 += k
	}
	//fmt.Println(sum1)
	fmt.Println(string(sum1 + 96))

	str2 := make([]string, 100)
	str2 = re.FindAllString(s2, -1)
	sum2 := 0
	for j := 0; j < len(str2); j++ {
		k, _ := strconv.Atoi(str2[j])
		sum2 += k
	}
	//fmt.Println(sum2)
	fmt.Println(string(sum2 + 96))

	//arr := make([]string, 100)
	//w1 := arr[0]
	//w2 := arr[]
	//len(w1) w1[1]

	//var i, num int
	//
	//for i, num := range str1 {
	//	var sum1 int
	//	sum1 = 0
	//	sum1 = sum1 + []int(str1[i])
	//}
}
