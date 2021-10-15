package entity

import (
	"time"

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

	MedicalTech1 := MedicalTech{
		Name:	"Boss",
		Email: "chalermkiet@gmail.com",
		Pass: "asdf",
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech1)

	MedicalTech2 := MedicalTech{
		Name:	"chalermkiet",
		Email: 	"boss@gmail.com",
		Pass: 	"sadf",
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech2)

	MedicalRecord1 := MedicalRecord{
		Hospital_Number: "123",
		Personal_ID:	"12341234",
		Patient_Name:	"masdf",
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