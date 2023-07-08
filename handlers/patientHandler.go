package handlers

import (
	"clinic/db/models"
	"clinic/initializers"

	"github.com/gin-gonic/gin"
)

var Patient struct {
	Name    string
	Age     int
	Gender  string
	Address string
	Phone   string
}

func CreatePatient(c *gin.Context) {
	c.Bind(&Patient)
	patient := models.Patient{
		Name:    Patient.Name,
		Age:     Patient.Age,
		Gender:  Patient.Gender,
		Address: Patient.Address,
		Phone:   Patient.Phone,
	}

	result := initializers.DB.Create(&patient)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"data": patient,
	})

}

func GetPatient(c *gin.Context) {
	var patients []models.Patient
	initializers.DB.Find(&patients)

	c.JSON(200, gin.H{
		"data": patients,
	})

}

func ShowPatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	initializers.DB.First(&patient, id)

	c.JSON(200, gin.H{
		"data": patient,
	})
}

func UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	c.Bind(&Patient)

	var patient models.Patient
	initializers.DB.First(&patient, id)

	initializers.DB.Model(&patient).Updates(models.Patient{
		Name:    Patient.Name,
		Age:     Patient.Age,
		Gender:  Patient.Gender,
		Address: Patient.Address,
		Phone:   Patient.Phone,
	})

	c.JSON(200, gin.H{
		"data": patient,
	})
}

func DeletePatient(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Patient{}, id)

	c.Status(200)
}
