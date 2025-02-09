package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter your text")

	hash := make(map[string]string)

	for {
		Input := bufio.NewReader(os.Stdin)
		UserInput, _ := Input.ReadString('\n')
		Splits := strings.Split(UserInput, " ")
		if UserInput == "Ping" {

			fmt.Println("Pong")
		}
		fmt.Println(Splits)
		fmt.Println(hash)
	}
}
