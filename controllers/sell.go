package controllers

import (
	"fmt"
	"net/http"

	database "BasicRestAPI/database"
	"BasicRestAPI/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Sell(c *gin.Context) {
	var sell map[string]interface{}

	if err := c.ShouldBindJSON(&sell); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingAccount models.Account
	err := database.DB.Where("id = ?", uint(sell["account_id"].(float64))).First(&existingAccount).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Cannot sell. Invalid account ID: %v", sell["account_id"]),
		})
		return
	}

	history := models.History{
		AccountID: uint(sell["account_id"].(float64)),
		Ticker:    sell["stock"].(string),
		Lot:       int(sell["lot"].(float64)),
		Type:      "sell",
	}

	tx := database.DB.Begin()

	result := tx.Create(&history)

	if result.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Order error",
			"error":   result.Error.Error(),
		})
		return
	}

	var existingPortfolio models.Portfolio
	err = tx.Where("account_id = ? AND ticker = ?", uint(sell["account_id"].(float64)), sell["stock"].(string)).First(&existingPortfolio).Error

	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Cannot sell %s. Portfolio for %s doesn't exist.", sell["stock"].(string), sell["stock"].(string)),
		})
		return
	}

	if float64(existingPortfolio.Lot) < sell["lot"].(float64) {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Cannot sell %s. Insufficient lot quantity.", sell["stock"].(string)),
		})
		return
	}

	updateResult := tx.Model(&models.Portfolio{}).
		Where("account_id = ? AND ticker = ?", uint(sell["account_id"].(float64)), sell["stock"].(string)).
		Update("lot", gorm.Expr("lot - ?", sell["lot"].(float64)))

	if updateResult.Error != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Order error",
			"error":   updateResult.Error.Error(),
		})
		return
	}

	if updateResult.RowsAffected > 0 && int(existingPortfolio.Lot)-int(sell["lot"].(float64)) == 0 {
		deleteResult := tx.Where("account_id = ? AND ticker = ?", uint(sell["account_id"].(float64)), sell["stock"].(string)).Delete(&models.Portfolio{})
		if deleteResult.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Order error",
				"error":   deleteResult.Error.Error(),
			})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "Your " + sell["stock"].(string) + " sell order is matched",
	})
}
