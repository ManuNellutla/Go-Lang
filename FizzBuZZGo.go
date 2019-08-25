package main

import (
	"fmt"
	"strconv"
)

func main() {

	for i := 1; i < 20; i++ {

		result := ""

		if i%3 == 0 {
			result += "fizz"
		}

		if i%5 == 0 {
			result += "buzz"
		}

		if (i%3 != 0) && (i%5 != 0) {
			result = strconv.Itoa(i)
		}

		fmt.Println(result)
	}
}
