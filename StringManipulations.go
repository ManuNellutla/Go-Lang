package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string

	fmt.Println("Enter list of words / sentence")

	_, err := fmt.Scanf("%q", &input)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("String entered : %v \n", input)
	fmt.Println("")
	words := strings.Split(input, " ")
	fmt.Printf("Number of words enterred %v \n", len(words))

	for i := range words {
		fmt.Println(words[i])
	}

	// substrings

	var x string = "this is a string."
	runes := []rune(x)
	fmt.Println(runes)
	fmt.Println(string(runes))

}
