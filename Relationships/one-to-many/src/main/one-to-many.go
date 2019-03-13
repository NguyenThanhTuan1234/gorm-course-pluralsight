package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
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
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	//db.Debug().Model(&Calendar{}).
	//	AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE") //Association between user and calendar will be deleted when delete a user

	db.Debug().Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
			Appointments: []Appointment{
				{Subject: "Spontaneous Whale Generation"},
				{Subject: "Saved from Vaccuum of Space"},
			},
		},
	})

	//u := User{}
	c := Calendar{}
	a := Appointment{}


	for i, _ := range c.Appointments {
		db.Find(&c).Related(&a)
		fmt.Println(&c.Appointments[i])
		fmt.Println(a)
	}

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
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	Subject		string
	Description	string
	StartTime	time.Time
	Length		uint
	CalendarID	uint
}