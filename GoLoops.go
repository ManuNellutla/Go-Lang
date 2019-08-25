package main

import "fmt"

func main() {

	n := 1

	for n < 10 {

		n *= 2
	}
	fmt.Println(n)

	for i := 0; i < 5; i++ {
		fmt.Println(n + i)
	}
}
