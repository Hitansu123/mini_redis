package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
) 


type Data struct{
	Keys string `gorm:"primarykey"`
	Values string
}
func Sqlite_setup(){
	
	db,err:=gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err!=nil{
		fmt.Println("error opening database")
	}
	db.AutoMigrate(&Data{})
	test:=Data{Keys: "test", Values: "test"}
	db.Create(test)

	



} 
