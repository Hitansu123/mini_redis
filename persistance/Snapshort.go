package persistance

import (
	"fmt"
	"time"
	"Building_Redis/secondaryDB"
	"gorm.io/gorm"
)
type record struct {
	Keys_data string 
	Values string
	TTL int
	ExpireAt time.Time
}

func Rdb_snapshort(db *gorm.DB) {
	//var Record []record
	
	ticker:=time.NewTicker(20*time.Second)
	defer ticker.Stop()

		for {
		select{
		case <-ticker.C:
			secondaryDB.DeleteFromDB()
			fmt.Println("Saving")
			secondaryDB.Store_SecondDB()	
		}
		//fmt.Println(Record)
		
	}
	
}
