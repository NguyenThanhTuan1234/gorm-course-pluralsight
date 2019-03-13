package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	db.DropTable(&User{})
	db.CreateTable(&User{})

	 u := User{
	 	FirstName: "Arthur",
	 	LastName: "Dent",
	 }
	 db.Debug().Create(&u)
	 fmt.Println(db.NewRecord(&u))
}

type User struct {
	gorm.Model
	FirstName	string
	LastName	string
}