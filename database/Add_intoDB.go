package database

import (
	"fmt"
	"strconv"
	"time"
)


func AddToDatabase(key, value ,ttl string){
	
	db:=Sqlite_setup()
	ttlInt,err:=strconv.ParseInt(ttl,10,64)
	if err!=nil{
		fmt.Println("Error in conversion",err)
	}
	min:=time.Now().Minute()
	exptime:=min+int(ttlInt)
	expireAt:=time.Now().Add(time.Duration(exptime))
	test:=Data{Keys_data: key, Values: value,TTL: int(ttlInt),ExpireAt: expireAt}
	result:=db.Create(&test)
	if result.Error!=nil{
		fmt.Println("Insert error")
	}
	fmt.Println("Inserted rows")


}
