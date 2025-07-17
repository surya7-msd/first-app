package main

import "fmt"

func main() {

	calc := func() int {
		return 4 + 10
	}

	fmt.Println("sum of two numbers : ", calc())
}
