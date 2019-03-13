package main

import (
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

	users := []User{
		{Username: "fprefect"},
		{Username: "tmacmillan"},
		{Username: "mrobot"},
	}

	for i := range users {
		db.Save(&users[i])
	}

	db.Debug().Save(&User{
		Username: "adent",
		Calendar: Calendar{
			Name: "Improbable Events",
			Appointments: []Appointment{
				{Subject: "Spontaneous Whale Generation", Attendees: users},
				{Subject: "Saved from Vaccuum of Space", Attendees: users},
			},
		},
	})



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
	Attendees	[]User `gorm:"many2many:appointment_user"`
}
