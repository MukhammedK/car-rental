package main

import (
	"car-rental-system/backend/config"
	"car-rental-system/backend/database"
	"car-rental-system/backend/internal/handlers"
	"car-rental-system/backend/internal/repositories"
	"car-rental-system/backend/internal/routes"
	"car-rental-system/backend/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	config.LoadEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	carRepo := repositories.NewCarRepository(db)
	carService := services.NewCarService(carRepo)
	carHandler := handlers.NewCarHandler(carService)

	userRepo := repositories.NewUserRepository(db)

	authService := services.NewAuthService(userRepo)

	authHandler := handlers.NewAuthHandler(authService)
	customerRepo := repositories.NewCustomerRepository(db)

	customerService := services.NewCustomerService(customerRepo)

	customerHandler := handlers.NewCustomerHandler(customerService)
	bookingRepo := repositories.NewBookingRepository(db)

	bookingService := services.NewBookingService(
		bookingRepo,
	)

	bookingHandler := handlers.NewBookingHandler(
		bookingService,
	)

	dashboardRepo := repositories.NewDashboardRepository(db)

	dashboardService := services.NewDashboardService(
		dashboardRepo,
	)

	dashboardHandler := handlers.NewDashboardHandler(
		dashboardService,
	)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5175",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
	}))

	routes.RegisterRoutes(
		router,
		carHandler,
		authHandler,
		customerHandler,
		bookingHandler,
		dashboardHandler,
	)
	router.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", nil)
	})

	router.GET("/cars", func(c *gin.Context) {
		c.HTML(200, "cars.html", nil)
	})

	router.GET("/customers", func(c *gin.Context) {
		c.HTML(200, "customers.html", nil)
	})
	router.Static("/static", "./static")

	router.Run(":8080")

	log.Println("PostgreSQL connected successfully")
}
