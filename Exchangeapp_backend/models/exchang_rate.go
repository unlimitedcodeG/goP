package models

import "time"

type ExchangeRate struct {
	ID           uint      `gorm:"primarykey" json:"_id"`
	FromCurrency string    `json:"fromCurrency" binding:"required"`
	ToCurrency   string    `json:"rate" binding:"required"`
	Rate         float64   `json:"rate" binding:"requirement"`
	Date         time.Time `json:"date"`
}
