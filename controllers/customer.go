package controllers

import (
	"dbo/entity"
	"dbo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindCustomers(c *gin.Context) {
	var input entity.GetCustomerReq

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customers, meta, err := models.GetListCustomers(input.Limit, input.Page, input.Search)
	fmt.Println(input.Limit)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers, "meta": meta})
}

func FindCustomerById(c *gin.Context) {
	var input entity.GetCustomerReq

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := models.GetDetailCustomer(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func CreateCustomer(c *gin.Context) {
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer := models.Customer{
		Name:    input.Name,
		Email:   input.Email,
		Address: input.Address,
	}
	savedCustomer, err := customer.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedCustomer})
}

func UpdateCustomer(c *gin.Context) {
	var input models.Customer
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer := models.Customer{
		ID:      input.ID,
		Name:    input.Name,
		Email:   input.Email,
		Address: input.Address,
	}
	savedCustomer, err := customer.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": savedCustomer})
}

func DeleteCustomer(c *gin.Context) {
	var input entity.GetCustomerReq

	err := c.ShouldBindUri(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		ID: input.ID,
	}
	_, err = customer.Delete()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
