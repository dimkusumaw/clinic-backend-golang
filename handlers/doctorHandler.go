package handlers

import (
	"clinic/db/models"
	"clinic/initializers"

	"github.com/gin-gonic/gin"
)

var Doctor struct {
	Name       string
	Gender     string
	Address    string
	Phone      string
	Specialist string
}

func CreateDoctor(c *gin.Context) {

	c.Bind(&Doctor)
	doctor := models.Doctor{
		Name:       Doctor.Name,
		Gender:     Doctor.Gender,
		Address:    Doctor.Address,
		Phone:      Doctor.Phone,
		Specialist: Doctor.Specialist,
	}

	result := initializers.DB.Create(&doctor)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": doctor,
	})
}

func GetDoctor(c *gin.Context) {
	var doctors []models.Doctor
	initializers.DB.Find(&doctors)

	c.JSON(200, gin.H{
		"data": doctors,
	})
}

func ShowDoctor(c *gin.Context) {
	id := c.Param("id")
	var doctor models.Doctor
	initializers.DB.First(&doctor, id)

	c.JSON(200, gin.H{
		"data": doctor,
	})
}

func UpdateDoctor(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&Doctor)

	var doctor models.Doctor
	initializers.DB.First(&doctor, id)

	initializers.DB.Model(&doctor).Updates(models.Doctor{
		Name:       Doctor.Name,
		Gender:     Doctor.Gender,
		Address:    Doctor.Address,
		Phone:      Doctor.Phone,
		Specialist: Doctor.Specialist,
	})

	c.JSON(200, gin.H{
		"data": doctor,
	})
}

func DeleteDoctor(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Doctor{}, id)

	c.Status(200)
}
