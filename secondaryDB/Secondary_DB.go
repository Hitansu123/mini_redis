package secondaryDB

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)
type record struct {
	Keys_data string 
	Values string
	TTL int
	ExpireAt time.Time
}


func Setup_secondDB() *gorm.DB{

	db,err:=gorm.Open(sqlite.Open("secondary_database.db"), &gorm.Config{})

	if err!=nil{
		fmt.Println("error opening database")
	}
	db.AutoMigrate(&record{})
	//go persistance.Rdb_snapshort(db)

	return db	


}
