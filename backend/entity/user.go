package entity

import(
	"time"
	"gorm.io/gorm"
)

type MedicalTech struct{
	gorm.Model
	Name	string
	Email	string		`gorm:"uniqueIndex"`
	Pass	string
	LabResults []LabResult `gorm:"foreignKey:MedicalTechID"`
}

type MedicalRecord struct{
	gorm.Model
	Hospital_Number	string
	Personal_ID		string		`gorm:"uniqueIndex"`
	Patient_Name	string
	Patient_Age		int
	Patient_Dob		time.Time
	Patient_Tel		string
	Register_Date 	time.Time
	LabResults []LabResult `gorm:"foreignKey:MedicalRecordID"`
}

type LabType struct{
	gorm.Model
	Name	string
	LabResults []LabResult `gorm:"foreignKey:LabTypeID"`
}

type LabRoom struct{
	gorm.Model
	Name	string
	Building	string
	floor	int
	LabResults []LabResult `gorm:"foreignKey:LabRoomID"`
}

type LabResult struct{
	gorm.Model
	
	MedicalTechID *uint
	MedicalTech		MedicalTech

	MedicalRecordID *uint
	MedicalRecord	MedicalRecord

	LabTypeID	*uint
	LabType	LabType

	Lab_Result	string
	Lab_Detail	string

	LabRoomID	*uint
	LabRoom	LabRoom

	AddedTime	time.Time
}