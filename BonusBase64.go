package main

import (
	"fmt"
	_ "strings"
)

func main() {
	choice := 0
	fmt.Println("Enter 1 to encode, 2 to decode")
	_, _ = fmt.Scan(&choice)
	if choice == 1 {
		fmt.Println(encode64())
	} else {
		fmt.Println(decode64())
	}
}

func encode64() string {
	m := make(map[uint8]string)
	m['A'] = "000000"
	m['B'] = "000001"
	m['C'] = "000010"
	m['D'] = "000011"
	m['E'] = "000100"
	m['F'] = "000101"
	m['G'] = "000110"
	m['H'] = "000111"
	m['I'] = "001000"
	m['J'] = "001001"
	m['K'] = "001010"
	m['L'] = "001011"
	m['M'] = "001100"
	m['N'] = "001101"
	m['O'] = "001110"
	m['P'] = "001111"
	m['Q'] = "010000"
	m['R'] = "010001"
	m['S'] = "010010"
	m['T'] = "010011"
	m['U'] = "010100"
	m['V'] = "010101"
	m['W'] = "010110"
	m['X'] = "010111"
	m['Y'] = "011000"
	m['Z'] = "011001"

	fmt.Println("Enter text to encode")
	toEncode := ""
	encoded := ""
	_, _ = fmt.Scan(&toEncode)
	for i := 0; i < len(toEncode); i++ {
		encoded = encoded + m[toEncode[i]]
	}
	return encoded
}

func decode64() string {
	m := make(map[string]string)
	m["000000"] = "A"
	m["000001"] = "B"
	m["000010"] = "C"
	m["000011"] = "D"
	m["000100"] = "E"
	m["000101"] = "F"
	m["000110"] = "G"
	m["000111"] = "H"
	m["001000"] = "I"
	m["001001"] = "J"
	m["001010"] = "K"
	m["001011"] = "L"
	m["001100"] = "M"
	m["001101"] = "N"
	m["001110"] = "O"
	m["001111"] = "P"
	m["010000"] = "Q"
	m["010001"] = "R"
	m["010010"] = "S"
	m["010011"] = "T"
	m["010100"] = "U"
	m["010101"] = "V"
	m["010110"] = "W"
	m["010111"] = "X"
	m["011000"] = "Y"
	m["011001"] = "Z"

	fmt.Println("Enter text to decode")
	toDecode := ""
	decoded := ""
	_, _ = fmt.Scan(&toDecode)
	for i := 0; i < len(toDecode); i = i + 6 {
		temp := toDecode[i : i+6]
		decoded = decoded + m[temp]
	}
	return decoded
}
