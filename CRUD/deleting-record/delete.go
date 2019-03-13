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

	db.Create(&User{
		FirstName: "Tricia",
		LastName: "Dent",
	})

	db.Create(&User{
		FirstName: "Arthur",
		LastName: "Dent",
	})

	db.Debug().Where("last_name LIKE ?", "Mac%").Delete(&User{})

	//db.Create(&u)
	//
	//db.Delete(&u)
	//
	//user := User{}
	//db.Debug().First(&user)
	//
	//fmt.Println(user)
}

type User struct {
	ID 			uint
	FirstName	string
	LastName	string
}
