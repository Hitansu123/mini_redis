package persistance

import (
	"fmt"
	"os"
)


func Store(Record []record){
	
	file,err:=os.Create("Record.json")
	if err!=nil{
		fmt.Println("Error",err)
	}
	for _,val:=range Record{
		data:=fmt.Sprintf("Keys=%v and Values=%v",val.Keys_data,val.Values)
		_,err:=file.WriteString(data)
		if err!=nil{
			fmt.Println("Error",err)
		}
	}
}
