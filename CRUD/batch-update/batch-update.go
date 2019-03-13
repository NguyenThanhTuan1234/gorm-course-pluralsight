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
		Salary: 50000,
	})

	db.Create(&User{
		FirstName: "Arthur",
		LastName: "Dent",
		Salary: 30000,
	})

	//db.Debug().Table("users").Where("last_name = ?", "Dent").Update("last_name", "MacMillan-Dent")
	db.Debug().Table("users").Where("salary > ?", 40000).Update("salary", gorm.Expr("salary + 5000"))
}

type User struct {
	gorm.Model
	FirstName	string
	LastName	string
	Salary		uint
}
