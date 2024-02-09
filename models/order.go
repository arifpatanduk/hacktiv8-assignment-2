package models

import "time"

type Order struct {
	ID           uint   `gorm:"primaryKey" json:"orderId"`
	CustomerName string `gorm:"not null;type:varchar(191)" json:"customerName"`
	Items []Item `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
	OrderedAt    time.Time `json:"orderedAt"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}