package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port= 5432 user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()


	db.DropTable(&User{})
	db.CreateTable(&User{})

	for _, user := range users {
		db.Create(&user)
	}

	//u := User{Username: "tmacmillan"}
	//db.Where(&u).First(&u)
	//fmt.Println(u)
	//
	//u.Lastname = "Beeblebrox"
	//db.Save(&u)
	//
	//user := User{}
	//db.Where(&u).First(&user)
	//fmt.Println(user)

	db.Where(&User{Username: "adent"}).Delete(&User{})
	fmt.Println("done")

	//u := User{}
	//db.Last(&u)
	//
	//fmt.Println(u)
}

type User struct {
	ID			uint
	Username	string
	Firstname	string
	Lastname	string
}

var users []User = []User{
	User{Username:"adent", Firstname: "Arthur", Lastname:"Dent"},
	User{Username:"fprefect", Firstname: "Ford", Lastname: "Prefect"},
	User{Username:"tmacmillan", Firstname: "Tricia", Lastname: "Macmillan"},
	User{Username:"mrobot", Firstname:"Marvin", Lastname:"Robot"},
}