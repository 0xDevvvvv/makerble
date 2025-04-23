package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/0xDevvvvv/makerble/config"
	"github.com/0xDevvvvv/makerble/internal/handlers"
	"github.com/0xDevvvvv/makerble/internal/middleware"
	"github.com/0xDevvvvv/makerble/pkg/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//load env variables
	config.LoadConfig()
	utils.InitJWT(config.AppConfig.JWTSecret)

	//set up db connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName)
	fmt.Println("Connecting with DSN:", psqlInfo)

	//open db connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database connection ", err)
	}
	defer db.Close()

	//run migrations
	config.RunMigrations(db)

	//set up handlers
	authHandler := handlers.NewAuthHandler(db) // set up auth handler with db so that it can interact with db

	//set up router
	router := gin.Default()

	//set up a public route for login and signup
	router.POST("/login", authHandler.Login)
	router.POST("/signup", authHandler.Signup)

	authorized := router.Group("/", middleware.AuthMiddleware())
	{

		authorized.GET("/patients", handlers.GetAllPatient)
		authorized.GET("/patients/:id", handlers.GetPatient)
		authorized.PUT("/patients", handlers.UpdatePatient)
		receptionist := authorized.Group("/", middleware.RoleMiddleware("receptionist"))
		{
			receptionist.POST("/patients", handlers.CreatePatient)
			receptionist.DELETE("/patients", handlers.DeletePatient)
		}

	}

	router.Run(":" + config.AppConfig.Port)
}
