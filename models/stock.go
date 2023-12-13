package models

type Stock struct {
	Ticker        string `gorm:"primaryKey;type:varchar(4)" json:"ticker"`
	LastPrice     int32  `gorm:"int" json:"last_price"`
	PreviousPrice int32  `gorm:"int" json:"previous_price"`
	OpenPrice     int32  `gorm:"int" json:"open_price"`
	Volume        int64  `gorm:"bigint" json:"volume"`
	Turnover      int64  `gorm:"bigint" json:"turnover"`
	Frequency     int64  `gorm:"bigint" json:"frequency"`
}
