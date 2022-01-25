package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("se-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Patientrecord{},
		&Prename{},
		&Gender{},
		&Bloodtype{},
		&Province{},
		&Personnel{},
	)

	db = database

	// Prename Data
	prename1 := Prename{
		Prename: "นาย",
	}
	db.Model(&Prename{}).Create(&prename1)

	prename2 := Prename{
		Prename: "นาง",
	}
	db.Model(&Prename{}).Create(&prename2)

	prename3 := Prename{
		Prename: "นางสาว",
	}
	db.Model(&Prename{}).Create(&prename3)

	// Gender Data
	gender1 := Gender{
		Gender: "ชาย",
	}
	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{
		Gender: "หญิง",
	}
	db.Model(&Gender{}).Create(&gender2)

	// BloodType Data
	bloodtype1 := Bloodtype{
		Bloodtype: "เอ",
	}
	db.Model(&Bloodtype{}).Create(&bloodtype1)

	bloodtype2 := Bloodtype{
		Bloodtype: "บี",
	}
	db.Model(&Bloodtype{}).Create(&bloodtype2)

	bloodtype3 := Bloodtype{
		Bloodtype: "เอบี",
	}
	db.Model(&Bloodtype{}).Create(&bloodtype3)

	bloodtype4 := Bloodtype{
		Bloodtype: "โอ",
	}
	db.Model(&Bloodtype{}).Create(&bloodtype4)

	// Province Data
	province1 := Province{
		Province: "นครราชสีมา",
	}
	db.Model(&Province{}).Create(&province1)

	province2 := Province{
		Province: "อุบลราชธานี",
	}
	db.Model(&Province{}).Create(&province2)

	// Personnel Data
	personnel1 := Personnel{
		Personnel: "ขยัน อดทด",
	}
	db.Model(&Personnel{}).Create(&personnel1)

	personnel2 := Personnel{
		Personnel: "สุภาพ อ่อนหวาน",
	}
	db.Model(&Personnel{}).Create(&personnel2)

	// Patientrecord 1
	db.Model(&Patientrecord{}).Create(&Patientrecord{
		Prename:        prename1,
		Firstname:      "นคร",
		Lastname:       "ศรีสรรณ์",
		Gender:         gender1,
		Idcardnumber:   "1234455678948",
		Age:            25,
		Birthday:       time.Now(),
		Bloodtype:      bloodtype3,
		Phonenumber:    "0855555555",
		Email:          "nakorn@test.com",
		Home:           "111 moo1",
		Province:       province1,
		Emergencyname:  "มาสาย ลาก่อน",
		Emergencyphone: "0111111111",
		Timestamp:      time.Now(),
		Personnel:      personnel1,
	})

}
