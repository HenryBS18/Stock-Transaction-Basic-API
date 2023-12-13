package models

type Account struct {
	ID         uint
	Name       string      `gorm:"type:varchar(100)" json:"name"`
	Portfolios []Portfolio `json:"portfolio"`
	Historys   []History   `json:"history"`
}
