package models

import "gorm.io/gorm"

type Medical struct {
	gorm.Model
	Patient    Patient
	Doctor     Doctor
	PatientID  int
	DoctorID   int
	Diseases   string `gorm: "type:text"`
	Medication string `gorm: "type:text"`
	Notes      string `gorm: "type:text"`
}
