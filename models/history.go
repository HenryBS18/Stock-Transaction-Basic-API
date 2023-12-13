package models

import (
	"time"
)

type History struct {
	AccountID        uint
	Ticker           string    `gorm:"type:varchar(4)" json:"ticker"`
	Price            int       `gorm:"int" json:"price"`
	Lot              int       `gorm:"int" json:"lot"`
	Type             string    `gorm:"type:varchar(4)" json:"type"`
	OrderTime        time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"order_time"`
	OrderMatchedTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"order_matched_time"`
}

type HistoryDTO struct {
	Stock            string    `json:"stock"`
	Price            int       `json:"price"`
	Lot              int       `json:"lot"`
	Type             string    `json:"type"`
	OrderTime        time.Time `json:"orderTime"`
	OrderMatchedTime time.Time `json:"orderMatchedTime"`
}

func (History) TableName() string {
	return "history"
}
