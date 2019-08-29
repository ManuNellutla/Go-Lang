package main

import "fmt"

func main() {

	a := 5
	b := 6
	fmt.Printf("\nvalues before swap a = %v and b= %v \n", a, b)
	a, b = b, a

	fmt.Printf("values after swap a = %v and b= %v", a, b)
}
