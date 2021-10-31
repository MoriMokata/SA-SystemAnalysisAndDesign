package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64-project.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(
		&MedicalRecord{}, &MedicalRecordOfficer{}, &NameTitle{}, &HealthInsurance{},
		&Nurse{}, &Disease{}, &Screening{},
		&Drug{}, &DrugAllergy{},
		&Doctor{}, &Department{}, &MedicalHistory{},
		&Hospital{}, &Refer{},
		&MedicalTech{}, &LabType{}, &LabRoom{}, &LabResult{},
	)

	db = database

	//setup MedicalRecord
	MedicalRecordOfficer1 := MedicalRecordOfficer{
		MedRecOfficer_Name:  "Rosie",
		MedRecOfficer_Email: "rosie@gmail.com",
		MedRecOfficer_Pass:  "111111a",
	}
	db.Model(&MedicalRecordOfficer{}).Create(&MedicalRecordOfficer1)

	MedicalRecordOfficer2 := MedicalRecordOfficer{
		MedRecOfficer_Name:  "Carla",
		MedRecOfficer_Email: "carla@gmail.com",
		MedRecOfficer_Pass:  "2222222a",
	}
	db.Model(&MedicalRecordOfficer{}).Create(&MedicalRecordOfficer2)

	NameTitle1 := NameTitle{
		Title: "นาง",
	}
	db.Model(&NameTitle{}).Create(&NameTitle1)

	NameTitle2 := NameTitle{
		Title: "นางสาว",
	}
	db.Model(&NameTitle{}).Create(&NameTitle2)

	NameTitle3 := NameTitle{
		Title: "นาย",
	}
	db.Model(&NameTitle{}).Create(&NameTitle3)

	NameTitle4 := NameTitle{
		Title: "เด็กหญิง",
	}
	db.Model(&NameTitle{}).Create(&NameTitle4)

	NameTitle5 := NameTitle{
		Title: "เด็กชาย",
	}
	db.Model(&NameTitle{}).Create(&NameTitle5)

	HealthInsurance1 := HealthInsurance{
		HealthInsurance_Name: "นักศึกษา",
		Detail:               "นักศึกษามหาวิทยาลัยเทคโนโลยีสุรนารีรักษาฟรี",
	}
	db.Model(&HealthInsurance{}).Create(&HealthInsurance1)

	HealthInsurance2 := HealthInsurance{
		HealthInsurance_Name: "บัตรทอง",
		Detail:               "สวัสดิการแห่งรัฐ 30 บาทรักษาทุกโรค",
	}
	db.Model(&HealthInsurance{}).Create(&HealthInsurance2)

	MedicalRecord1 := MedicalRecord{
		Hospital_Number: "1234",
		Personal_ID:     "567",

		Patient_Name:    "seeruk",
		Patient_Age:     23,
		Patient_dob:     time.Now(),
		Patient_Tel:     "094-XXXXXX",
		Register_Date:   time.Now(),
		HealthInsurance: HealthInsurance1,
		MedRecOfficer:   MedicalRecordOfficer1,
		NameTitle:       NameTitle2,
	}
	db.Model(MedicalRecord{}).Create(&MedicalRecord1)

	//setup Screening
	Disease1 := Disease{
		Name:        "sore throat",
		Description: "Most sore throats are caused by infections",
	}
	db.Model(&Disease{}).Create(&Disease1)

	password1, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	Nurse1 := Nurse{
		Name:  "nueng",
		Email: "nueng@gmail.com",
		Pass:  string(password1),
	}
	db.Model(&Nurse{}).Create(&Nurse1)

	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	Nurse2 := Nurse{
		Name:  "pakapon",
		Email: "pakapon@gmail.com",
		Pass:  string(password2),
	}
	db.Model(&Nurse{}).Create(&Nurse2)

	Screening1 := Screening{
		Symptoms:        "sore throat",
		Weight:          58,
		Height:          170,
		Temperature:     37,
		PulseRate:       120,
		RespirationRate: 80,
		MedRec:          MedicalRecord1,
		Disease:         Disease1,
		Nurse:           Nurse1,
	}
	db.Model(&Screening{}).Create(&Screening1)

	//setup DrugAllergy
	Drug1 := Drug{
		Drug_Name:       "Ruds1",
		Drug_properties: "RUCAAA",
		Drug_group:      "NRNS",
		Stock:           10,
	}
	db.Model(&Drug{}).Create(&Drug1)

	//setup refer
	/////
	///////
	//////////
	/////////////
	/////////////////

	//setup labresult
	Pass1, err := bcrypt.GenerateFromPassword([]byte("1"), 14)
	Pass2, err := bcrypt.GenerateFromPassword([]byte("123"), 14)

	MedicalTech1 := MedicalTech{
		Name:  "Boss",
		Email: "chalermkiet@gmail.com",
		Pass:  string(Pass1),
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech1)

	MedicalTech2 := MedicalTech{
		Name:  "chalermkiet",
		Email: "boss@gmail.com",
		Pass:  string(Pass2),
	}
	db.Model(&MedicalTech{}).Create(&MedicalTech2)

	LabType1 := LabType{
		Name: "blood",
	}
	db.Model(&LabType{}).Create(&LabType1)

	LabType2 := LabType{
		Name: "Atom",
	}
	db.Model(&LabType{}).Create(&LabType2)

	LabRoom1 := LabRoom{
		Name:     "room1",
		Building: "gf",
		floor:    1,
	}
	db.Model(&LabRoom{}).Create(&LabRoom1)

	LabResult1 := LabResult{
		MedicalTech:   MedicalTech1,
		MedicalRecord: MedicalRecord1,
		LabType:       LabType1,
		Lab_Result:    "wait MoKaTa",
		Lab_Detail:    "wait MokaTa",
		LabRoom:       LabRoom1,
		AddedTime:     time.Now(),
	}
	db.Model(&LabResult{}).Create(&LabResult1)
}
