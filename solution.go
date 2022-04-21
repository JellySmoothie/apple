package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args[1:]) != 2 {
		fmt.Println("Repeat input name!")
		return
	}

	// fmt.Println("Your name is : ", os.Args[1])
	fmt.Printf("full name is %s %s", os.Args[1], os.Args[2])

	// fmt.Println("Hello")
}
