package controllers

import (
	"assignment-2/database"
	"assignment-2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(ctx *gin.Context) {
	db := database.GetDB()
	var input models.Item

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orderInput := models.Item{}
	db.Create(&orderInput)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": orderInput,
	})
}