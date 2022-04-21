package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Repeat input name!")
		return
	}

	fmt.Println("Your name is : ", os.Args[1])

	// fmt.Println("Hello")
}
