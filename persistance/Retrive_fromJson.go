package persistance

import (
	"encoding/json"
	"fmt"
	"os"
)



func Retrive(){

	var tempdata map[string]string

	data,err:=os.ReadFile("Record.json")
	if err!=nil{
		fmt.Println("Cannot read the file",err)
	}
	json.Unmarshal(data,&tempdata)
	fmt.Println(tempdata)
}
