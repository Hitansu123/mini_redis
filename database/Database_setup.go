package database

import (
	"Building_Redis/persistance"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
) 


type Data struct{
	Keys_data string 
	Values string
}
func Sqlite_setup(){
	
	db,err:=gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err!=nil{
		fmt.Println("error opening database")
	}
	db.AutoMigrate(&Data{})
	//test:=Data{Keys_data: "test1", Values: "testvalue"}
	//result:=db.Create(&test)
	//if result.Error!=nil{
		//fmt.Println("Insert error")
	//}
	fmt.Println("Inserted rows")

	go persistance.Rdb_snapshort(db)

	



} 
