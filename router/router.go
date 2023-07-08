package router

import (
	"clinic/handlers"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.POST("/api/v1/register", handlers.Register)
	r.POST("/api/v1/login", handlers.Login)

	r.POST("/api/v1/doctor", handlers.CreateDoctor)
	r.GET("/api/v1/doctor", handlers.GetDoctor)
	r.GET("/api/v1/doctor/:id", handlers.ShowDoctor)
	r.PUT("/api/v1/doctor/:id", handlers.UpdateDoctor)
	r.DELETE("/api/v1/doctor/:id", handlers.DeleteDoctor)

	r.POST("/api/v1/patient", handlers.CreatePatient)
	r.GET("/api/v1/patient", handlers.GetPatient)
	r.GET("/api/v1/patient/:id", handlers.ShowPatient)
	r.PUT("/api/v1/patient/:id", handlers.UpdatePatient)
	r.DELETE("/api/v1/patient/:id", handlers.DeletePatient)

	r.Run()
}
