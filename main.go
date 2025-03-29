package main

import (
	"Building_Redis/database"
	"Building_Redis/persistance"
	"bufio"
	"fmt"
	"os"

	//"strconv"
	"strings"
	"sync"
	"time"
	"log"

	//"github.com/charmbracelet/log"
)

func main() {
	
	db:=database.Sqlite_setup()
	var wg sync.WaitGroup

	//erver := NewServer(":3000")
	//go func() { // Start the server in a separate goroutine
	//	if err := Server.Start(); err != nil {
	//fmt.Println("Server error:", err)
	//}
	//}()
	var hash sync.Map

	wg.Add(1)
	go persistance.Rdb_snapshort(db)
	wg.Wait()
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
	
	vals,_:=hash.Load("Keys_data")
	fmt.Println(vals)	
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
	if len(Splits)<4{
		log.Fatal("Please add Time to live after value")
	}
	hash.Store(Splits[1],Splits[2])
	database.AddToDatabase(Splits[1],Splits[2],Splits[3])
}
