package secondaryDB

import (
	"fmt"
	"time"
)


func GetData() []record{
	db:=Setup_secondDB()
	var secondDb []record		
	db.Raw("SELECT * from records").Scan(&secondDb)
	
	//fmt.Println("all data is",secondDb)
	var NotexpireRecord []record

	for _,vals:=range secondDb{
		if vals.ExpireAt.After(time.Now()){
			NotexpireRecord = append(NotexpireRecord,vals)
		}
	}


	fmt.Println("NotexpireRecord=",NotexpireRecord)

	return NotexpireRecord
}
