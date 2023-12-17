package controllers

import (
	"dbo/entity"
	"dbo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindOrders(c *gin.Context) {
	var input entity.GetCustomerReq

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, meta, err := models.GetListOrders(input.Limit, input.Page, input.Search)
	fmt.Println(input.Limit)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders, "meta": meta})
}

func FindOrderById(c *gin.Context) {
	var input entity.GetCustomerReq

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := models.GetDetailOrder(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func CreateOrder(c *gin.Context) {
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order := models.Order{
		TransactionNumber: input.TransactionNumber,
		TotalPrice:        input.TotalPrice,
		CreatedBy:         input.CreatedBy,
	}
	savedOrder, err := order.Save()

	for i, _ := range input.Details {
		input.Details[i].OrderID = savedOrder.ID
	}

	models.DetailSave(input.Details)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedOrder})
}
