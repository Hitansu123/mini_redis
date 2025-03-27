package persistance

import (
	"encoding/json"
	"fmt"
	"os"
)


func Store(Record []record){
	
	//file,err:=os.Create("Record.json")
	//if err!=nil{
		//fmt.Println("Error",err)
	//}
	file,err:=os.OpenFile("Record.json",os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("Cannot open the file",err)
	}
	for _,val:=range Record{
		data,err:=json.MarshalIndent(val,"","  ")
		if err!=nil{
			fmt.Println("Error marshal data",err)
		}
		_,err=file.Write(append(data,'\n'))
		if err!=nil{
			fmt.Println("Error in writing data",err)
		}
	}
}
