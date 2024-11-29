package main

import (
	"fmt"
	"strings"
)

func main() {
	b46encode("hello")
	// convert string to byte array
	// concatenate byte array
	// divide byte array into 6 bit chunks
	// convert six-bit bytes into eight-bit by prepending the prefix "00" in front of each group
	// convert each group from binary to decimal by finding its corresponding decimal value in the ASCII table
}

func b46encode(s string) {
	chars := strings.Split(s, "")
	binArray := []string{}
	for _, char := range chars {
		binArray = append(binArray, stringToBin(char))
	}

	fmt.Println(binArray)
}

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}
