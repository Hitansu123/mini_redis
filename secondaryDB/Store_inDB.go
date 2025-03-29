package secondaryDB

import (
	"Building_Redis/database"
	"fmt"
	"time"
)


func Store_SecondDB(){
	Alldata:=database.GetFromDatabase()
	
	db:=Setup_secondDB()
	if db==nil{
		fmt.Println("Failed to set up data")
	}
	var keyval string
	var values string
	var ttl int
	var expireat time.Time

	for _,val:=range Alldata{
		keyval=val.Keys_data
		values=val.Values
		ttl=val.TTL
		expireat=val.ExpireAt
		
		data:=record{Keys_data: keyval,Values: values,TTL: ttl,ExpireAt: expireat}

		result:=db.Create(&data)
		if result.Error !=nil{
			fmt.Println("Cannot insert into secondart db")
		}

	}
	fmt.Println("Insetion successful secondary db")	
}
