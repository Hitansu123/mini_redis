package database

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)


func AddToDatabase(key, value ,ttl string,wg *sync.WaitGroup){
	
	//defer wg.Done()
	db:=Sqlite_setup()
	ttlInt,err:=strconv.ParseInt(ttl,10,64)
	if err!=nil{
		fmt.Println("Error in conversion",err)
	}
	expireAt:=time.Now().Add(time.Duration(ttlInt)*time.Second)
	test:=Data{Keys_data: key, Values: value,TTL: int(ttlInt),ExpireAt: expireAt}
	result:=db.Create(&test)
	if result.Error!=nil{
		fmt.Println("Insert error")
	}
	fmt.Println("Inserted rows")


}
