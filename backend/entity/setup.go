package entity

import (
	"time"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB{
	return db
}

func SetupDatabase(){
	database, err := gorm.Open(sqlite.Open("LabResult"), &gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
	database.AutoMigrate(
		&MedicalTech{}, &MedicalRecord{}, &LabType{}, &LabRoom{}, &LabResult{},
	)
	db = database

	Pass1, err := bcrypt.GenerateFromPassword([]byte("1"), 14)
	Pass2, err := bcrypt.GenerateFromPassword([]byte("123"), 14)

	MedicalTech1 := MedicalTech{
		Name:	"Boss",
		Email: "chalermkiet@gmail.com",
		Pass: string(Pass1),
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech1)

	MedicalTech2 := MedicalTech{
		Name:	"chalermkiet",
		Email: 	"boss@gmail.com",
		Pass: 	string(Pass2),
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech2)

	MedicalRecord1 := MedicalRecord{
		Hospital_Number: "123",
		Personal_ID:	"12341234",
		Patient_Name:	"Somkiat kongkapan",
		Patient_Age:	12,
		Patient_Dob: time.Now(),
		Patient_Tel:	"0382175",
		Register_Date: time.Now(),
	}
	db.Model(&MedicalRecord{}).Create(&MedicalRecord1)

	

	LabType1:= LabType{
		Name: "blood",
	}
	db.Model(&LabType{}).Create(&LabType1)
	
	LabType2:= LabType{
		Name: "Atom",
	}
	db.Model(&LabType{}).Create(&LabType2)

	LabRoom1:= LabRoom{
		Name: "room1",
		Building: "gf",
		floor: 1,
	}
	db.Model(&LabRoom{}).Create(&LabRoom1)

}