package router

import (
	"os"
	"recycle/app/middlewares"
	userController "recycle/features/user/controller"
	userRepository "recycle/features/user/repository"
	userUsecase "recycle/features/user/usecase"

	rubbishController "recycle/features/rubbish/controller"
	rubbishRepository "recycle/features/rubbish/repository"
	rubbishUsecase "recycle/features/rubbish/usecase"

	locationController "recycle/features/location/controller"
	locationRepository "recycle/features/location/repository"
	locationUsecase "recycle/features/location/usecase"

	pickupController "recycle/features/pickup/controller"
	pickupRepository "recycle/features/pickup/repository"
	pickupUsecase "recycle/features/pickup/usecase"

	detailPickupController "recycle/features/detail_pickup/controller"
	detailPickupRepository "recycle/features/detail_pickup/repository"
	detailPickupUsecase "recycle/features/detail_pickup/usecase"

	aiUsecase "recycle/features/ai/usecase"
	aiController "recycle/features/ai/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// User
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userController := userController.NewUserControllers(userUsecase)

	// Rubbish
	rubbishRepository := rubbishRepository.NewRubbishRepository(db)
	rubbishUsecase := rubbishUsecase.NewRubbishUsecase(rubbishRepository)
	rubbishController := rubbishController.NewRubbishControllers(rubbishUsecase)

	// Location
	locationRepository := locationRepository.NewLocationRepository(db)
	locationUsecase := locationUsecase.NewLocationUsecase(locationRepository)
	locationController := locationController.NewLocationControllers(locationUsecase)

	// Pickup
	pickupRepository := pickupRepository.NewPickupRepository(db)
	pickupUsecase := pickupUsecase.NewPickupUsecase(pickupRepository)
	pickupController := pickupController.NewPickupControllers(pickupUsecase)

	// Detail Pickup
	detailPickupRepository := detailPickupRepository.NewDetailPickupRepository(db)
	detailPickupUsecase := detailPickupUsecase.NewDetailPickupUsecase(detailPickupRepository, rubbishRepository, pickupRepository, userRepository)
	detailPickupController := detailPickupController.NewPickupControllers(detailPickupUsecase, pickupUsecase, userUsecase)

	// AI
	godotenv.Load()
	openaiKey := os.Getenv("OPENAI_API_KEY")
	aiUsecase := aiUsecase.NewAIUsecase(nil, openaiKey)
	aiController := aiController.NewRubbishController(aiUsecase)

	// User & Admin CRUD
	user := e.Group("/users")
	user.POST("/register", userController.CreateUser)
	user.POST("/login", userController.Login)
	user.GET("", userController.GetAllUser, middlewares.JWTMiddleware())
	user.GET("/:id", userController.GetUser, middlewares.JWTMiddleware())
	user.PUT("/:id", userController.Update, middlewares.JWTMiddleware())
	user.DELETE("/:id", userController.Delete, middlewares.JWTMiddleware())
	e.GET("/verify", userController.VerifyAccount)

	rubbish := e.Group("/rubbish")
	rubbish.POST("", rubbishController.CreateRubbish, middlewares.JWTMiddleware())
	rubbish.GET("", rubbishController.GetAllRubbish, middlewares.JWTMiddleware())
	rubbish.GET("/type", rubbishController.GetRubbishByType, middlewares.JWTMiddleware())
	rubbish.GET("/:id", rubbishController.GetRubbish, middlewares.JWTMiddleware())
	rubbish.PUT("/:id", rubbishController.UpdateRubbish, middlewares.JWTMiddleware())
	rubbish.DELETE("/:id", rubbishController.DeleteRubbish, middlewares.JWTMiddleware())

	location := e.Group("/location")
	location.POST("", locationController.CreateLocation, middlewares.JWTMiddleware())
	location.GET("", locationController.GetAllLocation, middlewares.JWTMiddleware())
	location.GET("/city", locationController.GetLocationByCity, middlewares.JWTMiddleware())
	location.GET("/:id", locationController.GetLocation, middlewares.JWTMiddleware())
	location.PUT("/:id", locationController.UpdateLocation, middlewares.JWTMiddleware())
	location.DELETE("/:id", locationController.DeleteLocation, middlewares.JWTMiddleware())

	pickup := e.Group("/pickup")
	pickup.POST("", pickupController.CreatePickup, middlewares.JWTMiddleware())
	pickup.PUT("/:id", pickupController.UpdatePickup, middlewares.JWTMiddleware())
	pickup.DELETE("/:id", pickupController.DeletePickup, middlewares.JWTMiddleware())
	pickup.GET("", pickupController.GetAllPickup, middlewares.JWTMiddleware())
	pickup.GET("/status", pickupController.GetDataByStatus, middlewares.JWTMiddleware())

	detailPickup := e.Group("/detail/pickup")
	detailPickup.POST("", detailPickupController.CreateDetailPickup, middlewares.JWTMiddleware())
	detailPickup.GET("", detailPickupController.GetAllDetailPickup, middlewares.JWTMiddleware())

	e.POST("/type", aiController.GetRecyclableRecommendation)
}
