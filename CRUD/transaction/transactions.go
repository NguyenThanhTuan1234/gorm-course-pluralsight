package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port= 5432 user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})

	u := User {
		FirstName: "Marvin",
		LastName: "Robot",
	}

	tx := db.Begin()
	if err = tx.Create(&u).Error; err != nil{
		tx.Rollback()
	}

	u.LastName = "The Happy Robot"
	if err = tx.Save(&u).Error; err != nil {
		tx.Rollback()
	}
	tx.Commit()
 }


type User struct {
	ID 			uint
	FirstName	string
	LastName	string
}
