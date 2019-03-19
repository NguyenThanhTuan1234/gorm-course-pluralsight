package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	db, err := gorm.Open("postgres" , "user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	db.DropTableIfExists(&User{}, &Calendar{}, &Appointment{}, "appointment_user")
	db.AutoMigrate(&User{}, &Calendar{}, &Appointment{})

	user := User{
		UserName: "adent",
		FirstName: "Arthur",
		LastName: "Dent",
		Calendar: Calendar{Name: "Arthur's Calendar"},
	}
	fmt.Println("Creating")
	db.Create(&user)

	user.Calendar.Name = "Arthur's Itinerary"

	fmt.Println("Updating")
	db.Save(&user)


}

type User struct {
	gorm.Model
	UserName	string
	FirstName	string
	LastName	string
	Calendar	Calendar
}

func (u *User) BeforeSave() error {
	fmt.Println("Before Save")
	return nil
}

func (u *User) BeforeCreate() error {
	fmt.Println("Before Create")
	return nil
	//return errors.New("Can't create new user")
}

func (u *User) AfterSave() error {
	fmt.Println("After Save")
	return nil
}

func (u *User) AfterCreate() error {
	fmt.Println("After Create")
	return nil
}


type Calendar struct {
	gorm.Model
	Name	string
	UserID 	uint
	Appointments []*Appointment
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	CalendarID  uint
	Attendees   []*User `gorm:"many2many:appointment_user"`
}