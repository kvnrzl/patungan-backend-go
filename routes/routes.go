package routes

import (
	"bitbucket.org/bri_bootcamp/patungan-backend-go/middleware"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/models"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/pkg/db"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/controllers"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/repositories"
	"bitbucket.org/bri_bootcamp/patungan-backend-go/src/services"
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

	err = database.AutoMigrate(&models.User{}, &models.Category{}, &models.Campaign{}, &models.Donation{}, &models.Payment{})
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

	// ====================================================================

	campaignRepository := repositories.InitCampaignRepository()
	campaignService := services.InitCampaignService(database, campaignRepository)
	campaignController := controllers.InitCampaignController(campaignService, validate)

	campaignRoutes := e.Group("/api/v1")

	campaignRoutes.POST("/campaigns", middleware.AuthMiddleware(campaignController.Create))
	campaignRoutes.GET("/campaigns", campaignController.GetAll)
	//campaignRoutes.GET("/campaigns/:id", campaignController.GetByID)
	campaignRoutes.GET("/campaigns/:title", campaignController.GetByTitle)

	// ====================================================================

	categoryRepository := repositories.InitCategoryRepository()
	categoryService := services.InitCategoryService(database, categoryRepository)
	categoryController := controllers.InitCategoryController(categoryService, validate)

	categoryRoutes := e.Group("/api/v1")

	categoryRoutes.GET("/categories", categoryController.GetAll)
	categoryRoutes.GET("/categories/:id", categoryController.GetByID)

	// ====================================================================

	donationRepository := repositories.InitDonationRepository()
	donationService := services.InitDonationService(database, donationRepository, userRepository)
	donationController := controllers.InitDonationController(donationService, validate)

	donationRoutes := e.Group("/api/v1")

	donationRoutes.POST("/campaigns/:campaign_id/donate", middleware.AuthMiddleware(donationController.CreateDonationRegisteredUser))
	donationRoutes.POST("/campaigns/:campaign_id/donate/guest", donationController.CreateDonationGuestUser)

	// ====================================================================

	paymentRepository := repositories.InitPaymentRepository()
	paymentService := services.InitPaymentService(database, paymentRepository, donationRepository)
	paymentController := controllers.InitPaymentController(paymentService, validate)

	paymentRoutes := e.Group("/api/v1")

	paymentRoutes.POST("/payments", paymentController.CreatePayment)

	// ====================================================================
}
