package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=postgres dbname=gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	//********************Create data***********************//

	//db.DropTable(&User{})
	//db.CreateTable(&User{})
	//db.DropTable(&Calendar{})
	//db.CreateTable(&Calendar{})
	//db.DropTable(&Appointment{})
	//db.CreateTable(&Appointment{})
	//
	//users := map[string]*User{
	//	"adent":		&User{UserName:"adent", FirstName: "Arthur", LastName:"Dent"},
	//	"fprefect":		&User{UserName:"fprefect", FirstName: "Ford", LastName:"Prefect"},
	//	"tmacmillan":	&User{UserName:"tmacmillan", FirstName: "Tricia", LastName:"Macmillan"},
	//	"zbeeblebrox":	&User{UserName:"zbeeblebrox", FirstName: "Zaphod", LastName:"Beeblebrox"},
	//	"mrobot":		&User{UserName:"mrobot", FirstName: "Marvin", LastName:"Robot"},
	//}
	//for _, user := range users {
	//	user.Calendar = Calendar{Name: "Calendar"}
	//}
	//
	//users["adent"].AddAppointment(&Appointment{
	//	Subject:	"Save House",
	//	StartTime:	parseTime("1979-07-02 08:00"),
	//	Length:		60,
	//})
	//
	//users["fprefect"].AddAppointment(&Appointment{
	//	Subject:	"Get a Drink at Local Pub",
	//	StartTime:	parseTime("1979-07-02 10:00"),
	//	Length:		11,
	//	Attendees: 	[]*User{users["adent"]},
	//})
	//
	//users["fprefect"].AddAppointment(&Appointment{
	//	Subject:	"Hitch a ride",
	//	StartTime:	parseTime("1979-07-02 10:12"),
	//	Length:		60,
	//	Attendees: 	[]*User{users["adent"]},
	//})
	//
	//users["fprefect"].AddAppointment(&Appointment{
	//	Subject:	"Attend Poetry Reading",
	//	StartTime:	parseTime("1979-07-02 11:00"),
	//	Length:		30,
	//	Attendees: 	[]*User{users["adent"]},
	//})
	//
	//users["fprefect"].AddAppointment(&Appointment{
	//	Subject:	"Get Thrown into Space",
	//	StartTime:	parseTime("1979-07-02 10:40"),
	//	Length:		5,
	//	Attendees: 	[]*User{users["adent"]},
	//})
	//
	//users["fprefect"].AddAppointment(&Appointment{
	//	Subject:	"Get saved from Space",
	//	StartTime:	parseTime("1979-07-02 11:45"),
	//	Length:		1,
	//	Attendees: 	[]*User{users["adent"]},
	//})
	//
	//users["zbeeblebrox"].AddAppointment(&Appointment{
	//	Subject:	"Explore Planet Builder's HomeWorld",
	//	StartTime:	parseTime("1979-07-03 11:00"),
	//	Length:		240,
	//	Attendees: 	[]*User{users["adent"]},
	//})
	//
	//for _, user := range users {
	//	db.Save(&user)
	//}

	//*************************Retrieving Single Record******************************

	//u := User{}
	//db.Debug().First(&u)
	//db.Debug().FirstOrInit(&u, &User{UserName:"fprefect"})
	//db.Debug().FirstOrCreate(&u, &User{UserName:"lprosser"})
	//db.Debug().Last(&u)

	//*************************Retrieving Record Sets********************************

	//user1 := []User{}
	//db.Debug().Find(&users)
	//db.Debug().Find(&user1, &User{UserName: "fprefect"})
	//db.Debug().Find(&users, map[string]interface{}{"user_name":"fprefect"})
	//db.Debug().Find(&user1, "user_name = ?", "fprefect")
	//for _, u := range user1{
	//	fmt.Printf("\n%v\n", u)
	//}

	//**************************Where Clauses****************************************
	//users := []User {}
	////db.Debug().Where("user_name = ?", "adent").Find(&users)
	////db.Debug().Where(&User{UserName: "adent"}).Find(&users)
	////db.Debug().Where(map[string]interface{}{"user_name":"adent"}).Find(&users)
	////db.Debug().Where("user_name in (?)", []string{"adent", "tmacmillan"}).Find(&users)
	////db.Debug().Where("user_name like ?", "%mac%" ).Find(&users)
	////db.Debug().Where("user_name like ? and first_name = ? ", "%mac%", "Tricia" ).Find(&users)
	////db.Debug().Where("created_at < ?", time.Now()).Find(&users)
	////db.Debug().Where("created_at BETWEEN ? and ?", time.Now().Add(-30*24*time.Hour), time.Now()).Find(&users)
	////db.Debug().Not("user_name = ?", "adent").Find(&users)
	//db.Debug().Where("user_name = ?", "adent").Or("user_name = ?" , "fprefect").Find(&users)
	//for _, u := range users {
	//	fmt.Printf("\n%v\n", u)
	//}

	//***************************Preloading-Child-Objects******************************

	//users := []User{}
	//db.Debug().Preload("Calendar.Appointments").Find(&users)
	//for _, u := range users {
	//	fmt.Printf("\n%v\n", u.Calendar)
	//}

	//*****************************Limits-Offsets-Ordering******************************
	users := []User{}

	db.Debug().Limit(2).Offset(2).Order("first_name ").Find(&users)
	for _, u := range users {
		fmt.Printf("\n%v\n", u)
	}
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