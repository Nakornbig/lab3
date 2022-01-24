package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Age       uint8
	BirthDay  time.Time
}

type Patientrecord struct {
	gorm.Model

	Firstname      string
	Lastname       string
	Idcardnumber   string
	Age            uint8
	Birthday       time.Time
	Phonenumber    string
	Email          string
	Home           string
	Emergencyname  string
	Emergencyphone string
	Timestamp      time.Time

	//prename_id ทำหน้าที่เป็น FK
	PrenameID *uint
	Prename   Prename `gorm:"references:id"`

	//prename_id ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	//prename_id ทำหน้าที่เป็น FK
	BloodTypeID *uint
	BloodType   BloodType `gorm:"references:id"`

	//prename_id ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id"`

	//prename_id ทำหน้าที่เป็น FK
	PersonnelID *uint
	Personnel   Personnel `gorm:"references:id"`
}

type Prename struct {
	gorm.Model
	Prename       string
	Patientrecord []Patientrecord `gorm:"foreignKey:PrenameID"`
}

type Gender struct {
	gorm.Model
	Gender        string
	Patientrecord []Patientrecord `gorm:"foreignKey:GenderID"`
}

type BloodType struct {
	gorm.Model
	Bloodtype     string
	Patientrecord []Patientrecord `gorm:"foreignKey:BloodTypeID"`
}

type Province struct {
	gorm.Model
	Province      string
	Patientrecord []Patientrecord `gorm:"foreignKey:ProvinceID"`
}

type Personnel struct {
	gorm.Model
	Personnel     string
	Patientrecord []Patientrecord `gorm:"foreignKey:PersonnelID"`
}
