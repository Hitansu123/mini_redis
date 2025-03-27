package persistance

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)
type record struct {
	Keys_data string 
	Values string 
}

func Rdb_snapshort(db *gorm.DB) {
	var Record []record
	
	ticker:=time.NewTicker(5*time.Second)
	defer ticker.Stop()

	results:=db.Raw("SELECT * FROM data").Scan(&Record)
	//results:=db.Table("data").Find(&recordes)
	if results.Error!=nil{
		fmt.Println("Error in query",results.Error)
	}

	for {
		select{
		case <-ticker.C:
			fmt.Println("Saving")
			Store(Record)
		}
		
	}
	
}
