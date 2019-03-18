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

	//db.DropTableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
	//db.CreateTable(&User{}, &Calendar{}, &Appointment{}, &Attachment{})
	db.AutoMigrate(&User{}, &Calendar{}, &Appointment{}, &Attachment{})
	//db.Debug().Model(&User{}).ModifyColumn("first_name", "VARCHAR(100)")
	db.Debug().Exec(`ALTER TABLE "users" ALTER "first_name" TYPE VARCHAR(100)`)
}

type User struct {
	gorm.Model
	UserName	string
	FirstName	string		`sql:"type:VARCHAR(100)"`
	LastName	string
	Calendar	Calendar
}

type Calendar struct {
	gorm.Model
	Name	string
	UserID 	uint	`sql:"index:idx_calendar_user_id"`
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject				string
	Description 		string
	StartTime			time.Time
	Length				uint
	CalendarID			uint	`sql:"index:idx_appointment_calendar_id"`
	Reccuring			bool
	RecurrencePattern	string
	Attendees			[]*User `gorm:"many2many:appointment_user"`
	Attachments			[]Attachment
}

type Attachment struct {
	gorm.Model
	Data			[]byte
	AppointmentID	uint	`sql:"index:idx_attachment_appointment_id"`
}