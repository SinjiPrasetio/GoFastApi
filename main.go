package main

import (
	"GoFastApi/core"
	"GoFastApi/handler"
	"GoFastApi/middleware"
	"GoFastApi/migration"
	"GoFastApi/users"
	"GoFastApi/websocket"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbName := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	corsDomain := os.Getenv("APP_DOMAIN")

	fmt.Println("Connecting to the database...")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local&net_write_timeout=6000"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	migration.Migrate(db)

	if err != nil {
		log.Fatal(err.Error())
	}
	authService := core.NewAuthService()

	fmt.Println("Database connected")

	// Repository Declaration
	fmt.Println("Preparing Repository...")
	// =============
	userRepository := users.NewRepository(db)
	// =============
	fmt.Println("Repository loaded.")

	// Service Declaration
	fmt.Println("Preparing Service...")
	// =============
	userService := users.NewService(userRepository)
	// =============
	fmt.Println("Service loaded.")

	// Handler Declaration
	fmt.Println("Preparing Handler...")
	// =============
	userHandler := handler.NewUserHandler(userService, authService)
	// =============
	fmt.Println("Handler loaded.")

	// Defining router configuration
	fmt.Println("Defining router configuration...")
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://" + corsDomain},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	fmt.Println("Configuration loaded.")

	api := router.Group("/api/v1")

	// USERS
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.POST("/users/check-email", userHandler.CheckEmail)
	api.POST("/users/update-bio", middleware.Auth(authService, userService), userHandler.BioUpdateHandler)
	api.POST("/users/upload-avatar", middleware.Auth(authService, userService), userHandler.UploadAvatar)
	api.POST("/users/validate", middleware.Auth(authService, userService), userHandler.Validate)

	// Websocket initialize
	fmt.Println("Initiating Websocket Route")
	router.GET("/ws", websocket.WSHandler)
	fmt.Println("Websocket Deployed!")

	// Storage for public
	router.Static("/storage", "./public")

	// Starting Web Server
	fmt.Println("Starting Web Server...")
	router.Run(":8000")
}
