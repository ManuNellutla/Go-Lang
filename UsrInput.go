package main

import (
	"fmt"
)

func main() {
	//reading an integer
	var age int
	fmt.Println("What is your age?")

	_, err := fmt.Scan(&age)

	if err != nil {
		panic(err)
	}

	var naam string
	fmt.Println("What is your naam?")
	_, err = fmt.Scan(&naam)

	if err != nil {
		fmt.Println(err)
	}

	//reading a string
	//reader := bufio.NewReader(os.Stdin)
	//var name string
	//fmt.Println("What is your name?")
	//name, _ := reader.ReadString('\n')

	fmt.Println("Your name is ", naam, " and you are age ", age)
}
