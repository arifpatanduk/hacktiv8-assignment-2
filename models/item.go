package models

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `gorm:"not null;type:varchar(191)" json:"itemCode"`
	Description string `gorm:"not null;type:varchar(191)" json:"description"`
	Quantity    int    `gorm:"not null;type:int" json:"quantity"`
	OrderID     uint `json:"orderId"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}