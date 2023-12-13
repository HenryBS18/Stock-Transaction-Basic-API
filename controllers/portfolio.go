package controllers

import (
	database "BasicRestAPI/database"
	"BasicRestAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Portfolio(c *gin.Context) {
	var AccountID map[string]interface{}
	var portfolios []models.Portfolio
	var portfolioDTOs []models.PortfolioDTO

	if err := c.ShouldBindJSON(&AccountID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account models.Account
	err := database.DB.First(&account, "id = ?", AccountID["account_id"]).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invalid Client ID"})
		return
	}

	if err := database.DB.Where("account_id = ?", AccountID["account_id"]).Find(&portfolios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(portfolios) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No portfolios for the specified Client ID"})
		return
	}

	for _, portfolio := range portfolios {
		portfolioDTO := models.PortfolioDTO{
			Ticker:       portfolio.Ticker,
			Lot:          portfolio.Lot,
			AveragePrice: portfolio.AveragePrice,
		}
		portfolioDTOs = append(portfolioDTOs, portfolioDTO)
	}

	c.JSON(http.StatusOK, gin.H{
		"portfolio": portfolioDTOs,
	})
}
