package main

import (
	"clinic/db/models"
	"clinic/initializers"
)

func init() {
	initializers.SetupDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Doctor{}, &models.Patient{})
}
