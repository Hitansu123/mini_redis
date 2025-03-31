package database

import (
	"Building_Redis/models"
	"fmt"
	"time"

	//"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
) 


type Data struct{
	Keys_data string 
	Values string
	TTL int
	ExpireAt time.Time
}
func Sqlite_setup() *gorm.DB{
	
	db,err:=gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err!=nil{
		fmt.Println("error opening database")
	}
	db.AutoMigrate(&Data{},&models.Hash{},&models.Set{},&models.List{})
	//go persistance.Rdb_snapshort(db)

	return db	



} 
