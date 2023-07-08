package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	Name       string
	Gender     string
	Address    string
	Phone      string
	Specialist string
}
