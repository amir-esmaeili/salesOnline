package models

import "time"

type Product struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	SellerId  uint      `json:"seller_id"`
	Seller    User      `json:"seller" gorm:"OnDelete:CASCADE"`
	Price     uint      `json:"price"`
	Available uint      `json:"available"`
	Off       uint      `json:"off" gorm:"default:0"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}