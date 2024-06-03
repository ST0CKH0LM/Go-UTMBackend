package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	Products struct {
		Id              uint           `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
		Product_Id      string         `json:"product_id"`
		Product_name    string         `gorm:"not null" json:"product_name"`
		Product_type    string         `gorm:"not null" json:"product_type"`
		Product_price   int            `gorm:"not null" json:"product_price"`
		Product_brand   string         `json:"product_brand"`
		Product_detail  string         `json:"product_detail"`
		Product_desc    string         `json:"product_desc"`
		Product_count   int            `gorm:"not null" json:"product_count"`
		Product_sale    int            `json:"product_sale"`
		Product_gallery datatypes.JSON `json:"product_gallery"`
		Product_rate    int            `json:"rate"`
		Product_spac    string         `json:"product_spac"`
		Fav             bool           `json:"fav"`
		Merchant        string         `json:"merchant"`
		CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
		UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
		DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	}

	Number struct {
		Offsetnumber int
	}

//	TypeOfProduct struct {
//		id           uint
//		Product_type string `gorm:"not null"`
//		Thumbnail    string
//		createdAt    []uint8
//	}
)
