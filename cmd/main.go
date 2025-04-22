package main

import (
	"github.com/0xDevvvvv/makerble/config"
	"github.com/0xDevvvvv/makerble/internal/handlers"
	"github.com/0xDevvvvv/makerble/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	router := gin.Default()

	router.POST("/login", handlers.Login)

	authorized := router.Group("/", middleware.AuthMiddleware())
	{
		receptionist := authorized.Group("/", middleware.RoleMiddleware("receptionist"))
		{
			receptionist.GET("/patients", handlers.GetAllPatient)
			receptionist.GET("/patients/:id", handlers.GetPatient)
			receptionist.POST("/patients", handlers.CreatePatient)
			receptionist.PUT("/patients", handlers.UpdatePatient)
			receptionist.DELETE("/patients", handlers.DeletePatient)
		}
		doctor := router.Group("/", middleware.RoleMiddleware("doctor"))
		{
			doctor.GET("/patients", handlers.GetAllPatient)
			doctor.GET("/patients/:id", handlers.GetPatient)
			doctor.PUT("/patients", handlers.UpdatePatient)

		}
	}

	router.Run(":" + config.AppConfig.Port)
}
