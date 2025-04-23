package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/0xDevvvvv/makerble/config"
	"github.com/0xDevvvvv/makerble/internal/handlers"
	"github.com/0xDevvvvv/makerble/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	//load env variables
	config.LoadConfig()

	//set up db connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBUser,
		config.AppConfig.DBPassword, config.AppConfig.DBName)

	//open db connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database connection ", err)
	}
	defer db.Close()

	//run migrations
	config.RunMigrations(db)

	//set up router
	router := gin.Default()

	//set up a public route for login
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
