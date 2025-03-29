package secondaryDB

import "fmt"


func DeleteFromDB(){
	
	db:=Setup_secondDB()

	err:=db.Exec("DELETE FROM records;").Error

	if err!=nil{
		fmt.Println("Cannot delete data")
	}
	fmt.Println("Data deleted")

}
