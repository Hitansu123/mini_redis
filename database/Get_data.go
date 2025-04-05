package database

import "Building_Redis/models" 


func GetFromDatabase() []Data{
	db:=Sqlite_setup()
	var Alldata []Data
	
	db.Raw("SELECT * from data").Scan(&Alldata)
	
	//fmt.Println("all data is",Alldata)
	return Alldata
}

func GetFromListDatabase() []models.List{
	db:=Sqlite_setup()
	var Alldata []models.List
	
	db.Raw("SELECT * from lists").Scan(&Alldata)
	
	//fmt.Println("all data is",Alldata)
	return Alldata
}

