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
	//defer db.Close()

	db.DropTable(&User{})
	db.CreateTable(&User{})

	u := User {
		FirstName: "Ford",
		LastName: "Prefect",
	}

	db.Create(&u)

	fmt.Println(u)
	fmt.Println()

	//db.Debug().Model(&u).Update("first_name", "Zaphod")
	//u.FirstName = "Zaphod"
	//u.LastName = "Beeblebrox"
	//
	//db.Debug().Save(&u)

	db.Debug().Model(&u).Update(
		map[string]interface{}{
			"first_name" : "Zaphod",
			"last_name" : "Beeblebrox",
		})

	fmt.Println(u)

}

type User struct {
	gorm.Model
	FirstName	string
	LastName	string

}
