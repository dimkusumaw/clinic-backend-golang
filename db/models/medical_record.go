package models

import "gorm.io/gorm"

type Medical struct {
	gorm.Model
	Patient    Patient `gorm: "foreignKey:PatientID"`
	Doctor     Doctor  `gorm: "foreignKey:DoctorID"`
	PatientID  int
	DoctorID   int
	Diseases   string `gorm: "type:text"`
	Medication string `gorm: "type:text"`
	Notes      string `gorm: "type:text"`
}
