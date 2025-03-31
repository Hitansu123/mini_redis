package main

import (
	"Building_Redis/database"
	"Building_Redis/persistance"
	"Building_Redis/secondaryDB"
	"bufio"
	"fmt"

	"Building_Redis/implement_datastructure/lists"
	"os"

	//"strconv"
	"log"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
	//"github.com/charmbracelet/log"
)
func LoadData(hash *sync.Map,wg *sync.WaitGroup){
	//defer wg.Done()
	record:=secondaryDB.GetData()
	if record==nil{
		fmt.Println("Key does not exsist")
		return
	}
	for _,val:=range record{
		hash.Store(val.Keys_data,val.Values)
	}
	//fmt.Println("working")
}

func main() {
	
	db:=database.Sqlite_setup()

	database.DelExpireData()
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
			setvalue(&hash, Splits,&wg,db)
		case "GET":
			fmt.Println("Ok")
			GetKey(&hash, Splits,&wg)
		case "DEL":
			wg.Add(1)
			go func(){
				deleteData(Splits[1], &hash, &wg)
				wg.Done()
			}()
		case "EXPIRE":
			SetExpire(&hash, Splits, &wg)
		case "LPUSH":
			lists.LPush(Splits[1], Splits[2])
		case "LPOP":
			//LPop(Splits[1])
		case "RPUSH":
			lists.RPush(Splits[1], Splits[2])
		case "LRANGE":
			lists.LRange(Splits[1])
		}
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

func GetKey(hash *sync.Map, Splits []string,wg *sync.WaitGroup) {
	
	//fmt.Println("pk")
	wg.Add(1)
	go func() {
			go LoadData(hash,wg)
			wg.Done()
	}()
	key := Splits[1]
	

	value, ok := hash.Load(key)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Key does not exsist")
	}
	//wg.Wait()
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

func setvalue(hash *sync.Map, Splits []string,wg *sync.WaitGroup,db *gorm.DB) {
	if len(Splits)<4{
		log.Fatal("Please add Time to live after value")
	}
	hash.Store(Splits[1],Splits[2])
	
	wg.Add(2)
	go func(){
		defer wg.Done()
		persistance.Rdb_snapshort(db)
	}()
	go func(){
		defer wg.Done()
		database.AddToDatabase(Splits[1],Splits[2],Splits[3],wg)
	}()
}
