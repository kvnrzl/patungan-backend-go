package routes

import (
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/pkg/db"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/src/controllers"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/src/repositories"
	"bitbucket.org/bri_bootcamp/fp-patungan-backend-go/src/services"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var validate *validator.Validate

func SetupRoutes(e *echo.Echo) {

	validate = validator.New()

	// ====================================================================

	redisClient := db.ConnectToRedis()

	// ====================================================================

	database, err := db.OpenDB()
	if err != nil {
		logrus.Fatal("error connect to db")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		logrus.Fatal("error migrate table")
	}

	// ====================================================================
	userRepository := repositories.InitUserRepository()
	userService := services.InitUserService(database, redisClient, userRepository)
	userController := controllers.InitUserController(userService, validate)

	userRoutes := e.Group("/api/v1")

	userRoutes.POST("/login", userController.LoginHandler)
	userRoutes.POST("/register", userController.RegisterHandler)
	userRoutes.POST("/logout", userController.LogoutHandler)

}
