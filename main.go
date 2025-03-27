package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
	"Building_Redis/database"
)

func main() {
	
	database.Sqlite_setup()
	var wg sync.WaitGroup

	//erver := NewServer(":3000")
	//go func() { // Start the server in a separate goroutine
	//	if err := Server.Start(); err != nil {
	//fmt.Println("Server error:", err)
	//}
	//}()
	var hash sync.Map

	for {
		Input := bufio.NewReader(os.Stdin)
		UserInput, _ := Input.ReadString('\n')
		newInput := strings.TrimSpace(UserInput)
		Splits := strings.Split(newInput, " ")
		switch Splits[0] {
		case "Ping":
			fmt.Println("Pong")
		case "SET":
			setvalue(&hash, Splits)
		case "GET":
			GetKey(&hash, Splits)
		case "DEL":
			wg.Add(1)
			go deleteData(Splits[1], &hash, &wg)

		case "EXPIRE":
			SetExpire(&hash, Splits, &wg)
		}
		wg.Wait()
	}
}

func SetExpire(hash *sync.Map, Splits []string, wg *sync.WaitGroup) {
	key := Splits[1]
	_, ok := hash.Load(key)
	if ok {
		duration, _ := time.ParseDuration(Splits[2] + "s")
		//ExpirationTime := time.Now().Add(duration)
		time.AfterFunc(duration, func() {
			wg.Add(1)
			deleteData(key, hash, wg)
		})
	} else {
		fmt.Println("Key does not exsist")
	}
}

func GetKey(hash *sync.Map, Splits []string) {
	key := Splits[1]
	value, ok := hash.Load(key)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exsist")
	}
}

func deleteData(key string, hash *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	_, found := hash.Load(key)
	if found {
		hash.Delete(key)
	} else {
		fmt.Println("Key does not exsist")
	}
}

func setvalue(hash *sync.Map, Splits []string) {
	hash.Store(Splits[1], strings.Join(Splits[2:], " "))
}
