package controllers

import (
	"Taxi_service/configs"
	_ "Taxi_service/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()
	gin.SetMode(configs.AppSettings.AppParams.GinMode)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", PingPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}
	apiG := router.Group("/api", checkUserAuthentication)
	usersG := apiG.Group("/users")
	{
		usersG.POST("", CreateUsers)
		usersG.GET("", PrintUsers)
		usersG.GET("/:id", PrintUsersByID)
		usersG.PUT("/:id", EditUsers)
		usersG.PATCH("/:id", EditUsersRating)
		usersG.DELETE("/:id", DeleteUsers, BlockUsers)
	}

	ordersG := apiG.Group("/orders")
	{
		ordersG.POST("", CreateOrder)
		ordersG.GET("", GetAllOrders)
		ordersG.GET("/:id", GetAllOrdersByID)
		ordersG.PUT("/:id", UpdateOrderByID)
		ordersG.PATCH("/:id", ChecksOrderasResponse)
		ordersG.DELETE("/:id", DeleteOrderByID)
	}

	taxicompsG := apiG.Group("/taxicomps")
	{
		taxicompsG.POST("", CreateTaxicomp)
		taxicompsG.GET("", GetAllTaxiComp)
		taxicompsG.GET("/:id", GetAllTaxiCompByID)
		taxicompsG.PUT("/:id", UpdateTaxiCompByID)
		taxicompsG.DELETE("/:id", DeleteTaxiCompByID)
	}
	reportG := apiG.Group("/report")
	{
		reportG.GET("", Report)
	}
	distanceG := apiG.Group("/distances")
	{
		distanceG.PATCH("/:id", AddOrderDistanceandTotalByID)
	}
	return router
}
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
