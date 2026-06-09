package routes

import (
	"car-rental-system/backend/internal/handlers"
	"car-rental-system/backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	carHandler *handlers.CarHandler,
	authHandler *handlers.AuthHandler,
	customerHandler *handlers.CustomerHandler,
	bookingHandler *handlers.BookingHandler,
	dashboardHandler *handlers.DashboardHandler,
) {

	api := router.Group("/api")

	{
		cars := api.Group("/cars")

		cars.GET("", carHandler.GetAll)
		cars.GET("/:id", carHandler.GetByID)

		cars.POST("", carHandler.Create)

		cars.PUT("/:id", carHandler.Update)

		cars.DELETE("/:id", carHandler.Delete)
	}

	{
		auth := api.Group("/auth")

		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}
	{
		customers := api.Group("/customers")

		customers.Use(middleware.AuthMiddlewaring())

		customers.GET("", customerHandler.GetAll)
		customers.GET("/:id", customerHandler.GetByID)

		customers.POST("", customerHandler.Create)

		customers.PUT("/:id", customerHandler.Update)

		customers.DELETE("/:id", customerHandler.Delete)
	}
	{
		bookings := api.Group("/bookings")

		bookings.Use(middleware.AuthMiddlewaring())

		bookings.GET("", bookingHandler.GetAll)

		bookings.GET("/:id", bookingHandler.GetByID)

		bookings.POST("", bookingHandler.Create)
		bookings.PATCH(
			"/:id/status",
			bookingHandler.UpdateStatus,
		)
	}
	{
		dashboard := api.Group("/dashboard")

		dashboard.Use(
			middleware.AuthMiddlewaring(),
		)

		dashboard.GET(
			"",
			dashboardHandler.GetStats,
		)
	}
	{
		public := api.Group("/public")

		public.POST(
			"/customers",
			customerHandler.Create,
		)

		public.POST(
			"/bookings",
			bookingHandler.Create,
		)
	}
}
