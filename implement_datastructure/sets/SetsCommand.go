package sets

import (
	"Building_Redis/database"
	"Building_Redis/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

var db = database.Sqlite_setup()

func SAdd(set_key,value string){
		var first models.Set

	// Check if there is an existing first element
	result := db.Where("set_key = ?", set_key,value).First(&first)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
				// Insert new element at the new position
			setsItem := models.Set{
				SetKey:  set_key,
				Value:    value,
			}

			db.Create(&setsItem)
		} else {
			log.Println("Database error:", result.Error)
			return
		}
	}
	
			setsItem := models.Set{
				SetKey:  set_key,
				Value:    value,
			}

			db.Create(&setsItem)
}

func SDelete(set_key,values string){
	var first models.Set
	result := db.Where("set_key = ?", set_key).First(&first)
	if result.Error!=nil{
		fmt.Println("Key not found")
		return
	}
	fmt.Println(first.Value)
	db.Delete(&first)
}

func SRange(set_key string){
	var first []models.Set
	
	db.Table("sets").Where("set_key = ?", set_key).Find(&first)
	// Handle -1 case (fetch all)
	

	// Check for errors
	if len(first) == 0 {
		log.Println("No records found for list:", set_key)
		return
	}

	// Print results
	for _, val := range first {
		fmt.Println(val.Value)
	}
}
