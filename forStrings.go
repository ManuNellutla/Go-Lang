package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printout(msg string) {
	fmt.Println(msg)
}

func main() {
	for i := 0; i < 10; i++ {
		printout(strings.Join([]string{"Counter:", strconv.Itoa(i)}, ""))
	}
}
