package persistance

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)
type record struct {
	Keys_data string 
	Values string
	TTL int
	ExpireAt time.Time
}

func Rdb_snapshort(db *gorm.DB) {
	var Record []record
	
	ticker:=time.NewTicker(20*time.Second)
	defer ticker.Stop()

		for {
		select{
		case <-ticker.C:
			var newRecord []record
			results:=db.Raw("SELECT * FROM data").Scan(&newRecord)
	//results:=db.Table("data").Find(&recordes)
			if results.Error!=nil{
				fmt.Println("Error in query",results.Error)
				continue
			}
			Record=append(Record, newRecord...)

			//fmt.Println("Saving")
			Store(Record)
		}
		//fmt.Println(Record)
		
	}
	
}
