package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

func main(){
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	db.DropTableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
	db.CreateTable(&User{}, &Calendar{}, &Appointment{}, &Attachment{})

}

type User struct {
	gorm.Model
	UserName	string
	FirstName	string
	LastName	string
	Calendar	Calendar
}

type Calendar struct {
	gorm.Model
	Name	string
	UserID 	uint
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject		string
	Description string
	StartTime	time.Time
	Length		uint
	CalendarID	uint
	Attendees	[]*User `gorm:"many2many:appointment_user"`
	Attachments	[]Attachment
}

type Attachment struct {
	gorm.Model
	Data			[]byte
	AppointmentID	uint
}