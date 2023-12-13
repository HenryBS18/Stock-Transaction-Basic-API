package controllers

import (
	database "BasicRestAPI/database"
	"BasicRestAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Stock(c *gin.Context) {
	var stockTicker map[string]interface{}
	var stock models.Stock

	if err := c.ShouldBindJSON(&stockTicker); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("ticker = ?", stockTicker["stock"]).First(&stock).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid Stock Name"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"previousPrice": stock.PreviousPrice,
		"openPrice":     stock.OpenPrice,
		"lastPrice":     stock.LastPrice,
		"volume":        stock.Volume,
		"frequency":     stock.Frequency,
		"turnover":      stock.Turnover,
	})
}
