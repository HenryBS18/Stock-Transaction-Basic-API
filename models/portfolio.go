package models

type Portfolio struct {
	AccountID    uint
	Ticker       string  `gorm:"type:varchar(4)" json:"ticker"`
	Lot          int     `gorm:"int" json:"lot"`
	AveragePrice float64 `gorm:"float" json:"average_price"`
}

func (Portfolio) TableName() string {
	return "portfolio"
}

type PortfolioDTO struct {
	Ticker       string
	Lot          int
	AveragePrice float64
}
