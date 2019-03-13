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
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	u := User {
		FirstName: "Arthur",
		LastName: "Dent",
	}

	appointments := []Appointment{
		{Subject: "First"},
		{Subject: "Second", Attendees: []*User{&u}},
		{Subject: "Third"},
	}

	u.Appointments = appointments

	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u))

}

type User struct {
	gorm.Model
	FirstName	string
	LastName	string
	Appointments []Appointment
}

type Appointment struct {
	gorm.Model
	UserID		uint
	StartTime	*time.Time
	Duration	uint
	Attendees	[]*User
	Subject		string
	Description	string
}