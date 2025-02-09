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

	hash := make(map[string]string)
	var wg sync.WaitGroup

	for {
		Input := bufio.NewReader(os.Stdin)
		UserInput, _ := Input.ReadString('\n')
		Splits := strings.Split(UserInput, " ")
		if Splits[0] == "Ping" {

			fmt.Println("Pong")
		} else if Splits[0] == "SET" {
			wg.Add(1)
			go setvalue(&hash, Splits, &wg)
		}
	}
}

func setvalue(hash *map[string]string, Splits []string, wg *sync.WaitGroup) {
	defer wg.Done()
	(*hash)[Splits[1]] = strings.Join(Splits[2:], " ")
	wg.Wait()

}
