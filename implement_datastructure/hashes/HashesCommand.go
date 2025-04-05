package hashes

import (
	"Building_Redis/database"
	"Building_Redis/models"
	"fmt"

	"gorm.io/gorm"
)

var db=database.Sqlite_setup()

func HADD(hashkey,field,value string){
	var first models.Hash

	// Check if there is an existing first element
	result := db.Where("hash_key = ? AND field = ?", hashkey,field).First(&first)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
				// Insert new element at the new position
			hashItem := models.Hash{
				HashKey:  hashkey,
				Field: field,
				Value:    value,
			}
			db.Create(&hashItem)
		}
	}else{
			first.Value=value
			err:=db.Save(&first).Error
			if err!=nil{
				fmt.Println("Something is wrong in updating the value",err)
			}
		//fmt.Println(result.Error)
	}
}

func HGET(hashkey,field string){
	var first models.Hash

	result:=db.Where("hash_key = ? AND field = ?",hashkey,field).First(&first)
	
	if result.Error==nil{
		fmt.Println("value is ",first.Value)
	}else{
		fmt.Println("data not found",result.Error)
	}
}
