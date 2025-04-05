package secondaryDB

import "fmt"



func DeleteFromDB(){
	
	
	var db =Setup_secondDB()
	err:=db.Exec("DELETE FROM records;").Error

	if err!=nil{
		fmt.Println("Cannot delete data")
	}
	fmt.Println("Data deleted")

}

func DeleteSingleEle(keys string){
	var singledata record
	fmt.Println("nice")	
	db:=Setup_secondDB()
	data:=db.Table("records").Where("keys_data=?",keys).Delete(&singledata)
	if data.Error!=nil{
		fmt.Println("Cannot delete the data",data.Error)
	}
}



func DeleteFromList(){
	
	
	var db =Setup_secondDB()
	err:=db.Exec("DELETE FROM lists;").Error

	if err!=nil{
		fmt.Println("Cannot delete data")
	}
	fmt.Println("Data deleted")

}
