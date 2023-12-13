package models

import (
	"BasicRestAPI/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/stock_golang?parseTime=true"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.Account{})
	database.AutoMigrate(&models.Stock{})
	database.AutoMigrate(&models.Portfolio{})
	database.AutoMigrate(&models.History{})

	DB = database
}
