package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.GET("/orders/:orderId", controllers.GetOrder)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)
	router.PUT("/orders/:orderId", controllers.EditOrder)

	return router
}