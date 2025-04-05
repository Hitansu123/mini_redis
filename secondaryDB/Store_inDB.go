package secondaryDB

import (
	"Building_Redis/database"
	"Building_Redis/models"
	//Building_Redis/secondaryDB"
	"fmt"
	"time"
)
var db=Setup_secondDB()

func Store_SecondDB(){
	Alldata:=database.GetFromDatabase()
	
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

func Store_ListSecondDB(){
	AllListdata:=database.GetFromListDatabase()
	
	var list_key string 
	var values string
	var position int

	for _,val:=range AllListdata{
		list_key=val.ListKey
		values=val.Value
		position=val.Position
	
		listdata:=models.List{ListKey: list_key,Value: values,Position: position}

		result:=db.Create(&listdata)
		if result.Error!=nil{
			fmt.Println("Cannot insert list data") 
		}
	}
	fmt.Println("successful inserted list data")

}
