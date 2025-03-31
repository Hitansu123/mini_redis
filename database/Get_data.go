package database

import() 


func GetFromDatabase() []Data{
	db:=Sqlite_setup()
	var Alldata []Data
	
	db.Raw("SELECT * from data").Scan(&Alldata)
	
	//fmt.Println("all data is",Alldata)
	return Alldata
}

