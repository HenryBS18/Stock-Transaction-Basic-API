package controllers

import (
	"BasicRestAPI/models"
	"net/http"

	database "BasicRestAPI/database"

	"github.com/gin-gonic/gin"
)

func History(c *gin.Context) {
	var AccountID map[string]interface{}
	var history []models.History

	if err := c.ShouldBindJSON(&AccountID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAccount models.Account
	err := database.DB.Where("id = ?", AccountID["account_id"]).First(&existingAccount).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid account ID",
		})
		return
	}

	if err := database.DB.Where("account_id = ?", AccountID["account_id"]).Find(&history).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var historyDTOs []models.HistoryDTO
	for _, h := range history {
		historyDTO := models.HistoryDTO{
			Stock:            h.Ticker,
			Price:            h.Price,
			Lot:              h.Lot,
			Type:             h.Type,
			OrderTime:        h.OrderTime,
			OrderMatchedTime: h.OrderMatchedTime,
		}
		historyDTOs = append(historyDTOs, historyDTO)
	}

	c.JSON(http.StatusOK, gin.H{
		"history": historyDTOs,
	})
}
