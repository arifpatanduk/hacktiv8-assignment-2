package controllers

import (
	"assignment-2/database"
	"assignment-2/hooks"
	"assignment-2/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// struct for request structure and validation
var request struct {
	CustomerName string `json:"customerName"`
	Items        []struct {
		ItemCode    string `json:"itemCode" binding:"required"`
		Description string `json:"description" binding:"required"`
		Quantity    int    `json:"quantity" binding:"required"`
	} `json:"items" binding:"required,dive"`
}

func GetAllOrders(c *gin.Context) {
	// initialize the db
	db := database.GetDB()

	// get all orders from the database
	var orders []models.Order
	db.Preload("Items").Find(&orders)

	// map the response
	var response []gin.H
	for _, order := range orders {
		orderData := gin.H{
			"orderId":      order.ID,
			"customerName": order.CustomerName,
			"orderedAt":    order.OrderedAt,
			"items":        order.Items,
		}
		response = append(response, orderData)
	}

	c.JSON(http.StatusOK, hooks.SuccessResponse(response, "Successfully retrieved all order data"))
}

func GetOrder(c *gin.Context) {
	// initialize the db
	db := database.GetDB()

	orderID := c.Param("orderId")

	// validate that orderId is a valid uint
	parsedOrderID, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, hooks.ErrorResponse("Invalid orderId"))
		return
	}

	// get the order with the given orderId from the database
	var order models.Order
	if err := db.Preload("Items").First(&order, parsedOrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, hooks.ErrorResponse("Order not found"))
		return
	}

	// map the response
	response := gin.H{
		"orderId":      order.ID,
		"customerName": order.CustomerName,
		"orderedAt":    order.OrderedAt,
		"items":        order.Items,
	}

	c.JSON(http.StatusOK, hooks.SuccessResponse(response, "Successfully retrieved order data"))
}

func CreateOrder(c *gin.Context) {
	// initialize the db
	db := database.GetDB()
	
	// struct for request structure and validation
	// var request struct {
	// 	CustomerName string `json:"customerName" binding:"required"`
	// 	Items        []*struct {
	// 		ItemCode    string `json:"itemCode" binding:"required"`
	// 		Description string `json:"description" binding:"required"`
	// 		Quantity    int    `json:"quantity" binding:"required,min=1"`
	// 	} `json:"items" binding:"required,dive"`
	// }
	

	// validate the request must be a valid json
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, hooks.ErrorResponse(err.Error()))
		return
	}

	// create order data
	order := models.Order{
		CustomerName: request.CustomerName,
		OrderedAt:    time.Now(),
		Items:        []models.Item{},
	}

	// add items to the order
	for _, itemReq := range request.Items {
		item := models.Item{
			ItemCode:    itemReq.ItemCode,
			Description: itemReq.Description,
			Quantity:    itemReq.Quantity,
		}
		order.Items = append(order.Items, item)
	}

	// store the order to db
	err := db.Create(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, hooks.ErrorResponse("Error creating order data"))
		return
	}

	// map the response
	c.JSON(http.StatusOK, hooks.SuccessResponse(gin.H{
		"orderId":    order.ID,
		"customerName": order.CustomerName,
		"orderedAt":  order.OrderedAt,
		"items":      order.Items,
	}, "Successfully created the order data"))
}

func DeleteOrder(c *gin.Context) {
	// initialize the db
	db := database.GetDB()

	orderID := c.Param("orderId")

	// validate that orderId is a valid uint
	parsedOrderID, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, hooks.ErrorResponse("Invalid orderId"))
		return
	}

	// cek if the order exists
	var existingOrder models.Order
	if err := db.First(&existingOrder, parsedOrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, hooks.ErrorResponse("Order not found"))
		return
	}

	// delete all the child items
	for _, item := range existingOrder.Items {
		if err := db.Delete(&item).Error; err != nil {
			c.JSON(http.StatusInternalServerError, hooks.ErrorResponse("Failed to delete items"))
			return
		}
	}

	// delete the order from the database
	if err := db.Delete(&existingOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, hooks.ErrorResponse("Failed to delete order"))
		return
	}

	// return response
	c.JSON(http.StatusOK, hooks.SuccessResponse(nil, "Successfully deleted order data for orderID: " + orderID))
}


func EditOrder(c *gin.Context) {
	// initialize the db
	db := database.GetDB()

	orderID := c.Param("orderId")

	// validate that orderId is a valid uint
	parsedOrderID, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, hooks.ErrorResponse("Invalid orderId"))
		return
	}

	// check if the order exists
	var existingOrder models.Order
	if err := db.Preload("Items").First(&existingOrder, parsedOrderID).Error; err != nil {
		c.JSON(http.StatusNotFound, hooks.ErrorResponse("Order not found"))
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, hooks.ErrorResponse(err.Error()))
		return
	}
	
	// delete existing item
	// acctually it could be simpler if there was a master table for items/products (just check the itemCode or productID, then update the quantity)
	db.Where("order_id = ?", parsedOrderID).Delete(&models.Item{})

	// add new items to the order
	var newItems []models.Item
	for _, newItem := range request.Items {
		item := models.Item{
			ItemCode:    newItem.ItemCode,
			Description: newItem.Description,
			Quantity:    newItem.Quantity,
			OrderID:     uint(parsedOrderID),
		}
		newItems = append(newItems, item)
	}

	// update the order data
	existingOrder.CustomerName = request.CustomerName
	existingOrder.Items = newItems
	db.Save(&existingOrder)

	// map the response
	response := gin.H{
		"orderId":      existingOrder.ID,
		"customerName": existingOrder.CustomerName,
		"orderedAt":    existingOrder.OrderedAt,
		"items":        existingOrder.Items,
	}

	c.JSON(http.StatusOK, hooks.SuccessResponse(response, "Successfully updated order data"))
}

