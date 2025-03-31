package secondaryDB

import "fmt"


func GetData() []record{
	db:=Setup_secondDB()
	var secondDb []record		
	db.Raw("SELECT * from records").Scan(&secondDb)
	
	fmt.Println("all data is",secondDb)
	return secondDb
}
