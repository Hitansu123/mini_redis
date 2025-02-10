package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter your text")

	var hash sync.Map
	//var wg sync.WaitGroup

	for {
		Input := bufio.NewReader(os.Stdin)
		UserInput, _ := Input.ReadString('\n')
		newInput := strings.TrimSpace(UserInput)
		Splits := strings.Split(newInput, " ")
		if Splits[0] == "Ping" {

			fmt.Println("Pong")
		} else if Splits[0] == "SET" {
			setvalue(&hash, Splits)
		} else if Splits[0] == "GET" {
			key := Splits[1]
			value, ok := hash.Load(key)
			if ok {
				fmt.Println(value)
			} else {
				fmt.Println("Key does not exsist")
			}
		}
	}
}

func setvalue(hash *sync.Map, Splits []string) {
	hash.Store(Splits[1], strings.Join(Splits[2:], " "))
}
