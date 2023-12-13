package controllers

import (
	"fmt"
	"net/http"

	database "BasicRestAPI/database"
	"BasicRestAPI/models"

	"github.com/gin-gonic/gin"
)

func Buy(c *gin.Context) {
	var buy map[string]interface{}
	var stockPortfolio models.Portfolio

	if err := c.ShouldBindJSON(&buy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(buy)

	var existingAccount models.Account
	err := database.DB.Where("id = ?", uint(buy["account_id"].(float64))).First(&existingAccount).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Cannot buy. Invalid account ID: %v", buy["account_id"]),
		})
		return
	}

	var existingStock models.Stock
	err = database.DB.Where("ticker = ?", buy["stock"].(string)).First(&existingStock).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Cannot buy %s. Stock doesn't exist.", buy["stock"].(string)),
		})
		return
	}

	history := models.History{
		AccountID: uint(buy["account_id"].(float64)),
		Ticker:    buy["stock"].(string),
		Price:     int(buy["price"].(float64)),
		Lot:       int(buy["lot"].(float64)),
		Type:      "buy",
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

	err = tx.Where("account_id = ? AND ticker = ?", uint(buy["account_id"].(float64)), buy["stock"].(string)).First(&stockPortfolio).Error

	if err != nil {
		portfolio := models.Portfolio{
			AccountID:    uint(buy["account_id"].(float64)),
			Ticker:       buy["stock"].(string),
			Lot:          int(buy["lot"].(float64)),
			AveragePrice: buy["price"].(float64),
		}
		result = tx.Create(&portfolio)

		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Order error",
				"error":   result.Error.Error(),
			})
			return
		}
	} else {
		newAveragePrice := (stockPortfolio.AveragePrice*float64(stockPortfolio.Lot) + buy["price"].(float64)*float64(buy["lot"].(float64))) / float64(stockPortfolio.Lot+int(buy["lot"].(float64)))

		result = tx.Model(&stockPortfolio).
			Where("account_id = ? AND ticker = ?", uint(buy["account_id"].(float64)), buy["stock"].(string)).
			Updates(map[string]interface{}{
				"lot":           stockPortfolio.Lot + int(buy["lot"].(float64)),
				"average_price": newAveragePrice,
			})

		if result.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Order error",
				"error":   result.Error.Error(),
			})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"message": "Your " + buy["stock"].(string) + " buy order is matched",
	})
}
