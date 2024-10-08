package controller

import (
	"net/http"

	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/models"
	"github.com/felipemacedo1/go-antifraud-ms/pkg/app/service"

	"github.com/gin-gonic/gin"
)

type FraudController struct {
	Service *service.CheckFraudService
}

func NewFraudController(service *service.CheckFraudService) *FraudController {
	return &FraudController{
		Service: service,
	}
}

func (fc *FraudController) CheckFraudHandler(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := fc.Service.Execute(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
