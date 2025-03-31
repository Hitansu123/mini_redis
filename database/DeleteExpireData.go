package database

import (
	"fmt"
	"time"
)

func DelExpireData(){
	db:=Sqlite_setup()
	
	verify:=db.Where("expire_at <= ?",time.Now()).Delete(&Data{})

	if verify.Error!=nil{
		fmt.Println("Can not Delete expire data")
	}
}
