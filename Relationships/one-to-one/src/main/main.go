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
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})

	db.Debug().Model(&Calendar{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE") //Association between user and calendar will be deleted when delete a user

	db.Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
		},
	})

	//u := User{}
	//c := Calendar{}
	//db.Find(&u).Related(&c, "calendar")
	//fmt.Println(u)
	//fmt.Println()
	//fmt.Println(c)

}

type User struct {
	gorm.Model
	Username	string
	FirstName	string
	LastName	string
	Calendar	Calendar
}

type Calendar struct {
	gorm.Model
	Name	string
	UserID	uint
}
