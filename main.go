package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter your text")

	var UserInput string
	for {
		fmt.Scan(&UserInput)
		if UserInput == "Ping" {

			fmt.Println("Pong")
		}
	}
}
