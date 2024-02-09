package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"

	_ "assignment-2/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order API
// @version 1.0
// @description Simple REST API for order with items
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.GET("/orders/:orderId", controllers.GetOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)
	router.PUT("/orders/:orderId", controllers.EditOrder)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}