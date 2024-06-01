package serializers

import (
	"time"

	"gitlab.com/Std217/test/model"
	"gorm.io/gorm"
)

type ProductResponse struct {
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
	Product_gallery string         `json:"product_gallery"`
	Product_rate    int            `json:"rate"`
	Product_spac    string         `json:"product_spac"`
	Fav             bool           `json:"fav"`
	Merchant        string         `json:"merchant"`
	CreatedAt       string         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       string         `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

func NewProductResponse(product model.Products) ProductResponse {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		loc = time.UTC
	}
	return ProductResponse{
		Id:              product.Id,
		Product_Id:      product.Product_Id,
		Product_name:    product.Product_name,
		Product_type:    product.Product_type,
		Product_price:   product.Product_price,
		Product_brand:   product.Product_brand,
		Product_detail:  product.Product_detail,
		Product_count:   product.Product_count,
		Product_gallery: product.Product_gallery,
		Product_rate:    product.Product_rate,
		Product_spac:    product.Product_spac,
		Fav:             product.Fav,
		Merchant:        product.Merchant,
		CreatedAt:       time.Time(product.CreatedAt).In(loc).Format(time.RFC3339),
		UpdatedAt:       time.Time(product.UpdatedAt).In(loc).Format(time.RFC3339),
	}
}
