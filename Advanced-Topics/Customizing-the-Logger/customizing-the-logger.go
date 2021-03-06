package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	db.LogMode(true)

	seedDB(db)
	db.Find(&User{})


}

func LongMeetings(db *gorm.DB) *gorm.DB {
	return db.Where("length > ?", 60)
}

func seedDB(db *gorm.DB){

	db.DropTable(&User{})
	db.CreateTable(&User{})
	db.DropTable(&Calendar{})
	db.CreateTable(&Calendar{})
	db.DropTable(&Appointment{})
	db.CreateTable(&Appointment{})

	users := map[string]*User{
		"adent":		&User{UserName:"adent", FirstName: "Arthur", LastName:"Dent"},
		"fprefect":		&User{UserName:"fprefect", FirstName: "Ford", LastName:"Prefect"},
		"tmacmillan":	&User{UserName:"tmacmillan", FirstName: "Tricia", LastName:"Macmillan"},
		"zbeeblebrox":	&User{UserName:"zbeeblebrox", FirstName: "Zaphod", LastName:"Beeblebrox"},
		"mrobot":		&User{UserName:"mrobot", FirstName: "Marvin", LastName:"Robot"},
	}
	for _, user := range users {
		user.Calendar = Calendar{Name: "Calendar"}
	}

	users["adent"].AddAppointment(&Appointment{
		Subject:	"Save House",
		StartTime:	parseTime("1979-07-02 08:00"),
		Length:		60,
	})

	users["fprefect"].AddAppointment(&Appointment{
		Subject:	"Get a Drink at Local Pub",
		StartTime:	parseTime("1979-07-02 10:00"),
		Length:		11,
		Attendees: 	[]*User{users["adent"]},
	})

	users["fprefect"].AddAppointment(&Appointment{
		Subject:	"Hitch a ride",
		StartTime:	parseTime("1979-07-02 10:12"),
		Length:		60,
		Attendees: 	[]*User{users["adent"]},
	})

	users["fprefect"].AddAppointment(&Appointment{
		Subject:	"Attend Poetry Reading",
		StartTime:	parseTime("1979-07-02 11:00"),
		Length:		30,
		Attendees: 	[]*User{users["adent"]},
	})

	users["fprefect"].AddAppointment(&Appointment{
		Subject:	"Get Thrown into Space",
		StartTime:	parseTime("1979-07-02 10:40"),
		Length:		5,
		Attendees: 	[]*User{users["adent"]},
	})

	users["fprefect"].AddAppointment(&Appointment{
		Subject:	"Get saved from Space",
		StartTime:	parseTime("1979-07-02 11:45"),
		Length:		1,
		Attendees: 	[]*User{users["adent"]},
	})

	users["zbeeblebrox"].AddAppointment(&Appointment{
		Subject:	"Explore Planet Builder's HomeWorld",
		StartTime:	parseTime("1979-07-03 11:00"),
		Length:		240,
		Attendees: 	[]*User{users["adent"]},
	})

	for _, user := range users {
		db.Save(&user)
	}
}

type UserViewModel struct {
	FirstName	string
	LastName	string
	CalendarName string `gorm:"column:name"`
}

func parseTime(timeRaw string) time.Time {
	const timeLayout = "2006-01-02 15:04"
	t, _ := time.Parse(timeLayout, timeRaw)
	return t
}

type User struct {
	gorm.Model
	UserName	string
	FirstName	string
	LastName	string
	Calendar	Calendar
}

func(u *User) AddAppointment(appt *Appointment) {
	u.Calendar.Appointments = append(u.Calendar.Appointments, appt)
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
}