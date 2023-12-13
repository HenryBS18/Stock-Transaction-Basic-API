package main

import (
	"net/http"

	"BasicRestAPI/controllers"
	db "BasicRestAPI/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDatabase()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/stock", controllers.Stock)
	r.GET("/portfolio", controllers.Portfolio)
	r.GET("/history", controllers.History)
	r.POST("/buy", controllers.Buy)
	r.POST("/sell", controllers.Sell)

	r.Run()
}
