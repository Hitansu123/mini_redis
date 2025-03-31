package lists

import (
	"Building_Redis/database"
	"Building_Redis/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// Use a single database connection
var db = database.Sqlite_setup()

func LPush(listkey, values string) {
	var first models.List
	var newpos int

	// Check if there is an existing first element
	result := db.Where("list_key = ?", listkey).Order("position ASC").First(&first)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// If no record exists, start position from 0
			newpos = 0
		} else {
			log.Println("Database error:", result.Error)
			return
		}
	} else {
		// Assign new position
		newpos = first.Position - 1
	}

	// Insert new element at the new position
	listItem := models.List{
		ListKey:  listkey,
		Value:    values,
		Position: newpos,
	}

	db.Create(&listItem)
}

func RPush(listkey, values string) {
	// Implement RPush logic
	var first models.List
	var newpos int

	// Check if there is an existing first element
	result := db.Where("list_key = ?", listkey).Order("position DESC").First(&first)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// If no record exists, start position from 0
			newpos = 0
		} else {
			log.Println("Database error:", result.Error)
			return
		}
	} else {
		// Assign new position
		newpos = first.Position +1
	}

	// Insert new element at the new position
	listItem := models.List{
		ListKey:  listkey,
		Value:    values,
		Position: newpos,
	}

	db.Create(&listItem)
}

func LRange(listkey string) {
	var first []models.List
	
	db.Table("lists").Where("list_key = ?", listkey).Find(&first)
	// Handle -1 case (fetch all)
	

	// Check for errors
	if len(first) == 0 {
		log.Println("No records found for list:", listkey)
		return
	}

	// Print results
	for _, val := range first {
		fmt.Println(val.Value)
	}	
}


