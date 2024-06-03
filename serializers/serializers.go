package serializers

import (
	"time"

	"gitlab.com/Std217/test/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	ProductResponse struct {
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
		CreatedAt       string         `gorm:"column:created_at;AUTO_INCREMENT" json:"created_at"`
		UpdatedAt       string         `gorm:"column:updated_at;AUTO_INCREMENT" json:"updated_at"`
		DeletedAt       gorm.DeletedAt `gorm:"AUTO_INCREMENT" json:"deleted_at"`
	}

	UsersResponse struct {
		Id          int            `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
		Name        string         `gorm:"not null" json:"name"`
		SName       string         `json:"surname"`
		Username    string         `gorm:"not null" json:"username"`
		Password    string         `gorm:"not null" json:"password"`
		Email       string         `gorm:"not null" json:"email"`
		Phone       string         `gorm:"not null" json:"phone"`
		DateOfBirth string         `json:"dateofbirth"`
		Gender      string         `json:"gender"`
		Thumbnail   string         `json:"thumbnail"`
		CreatedAt   string         `gorm:"column:created_at;AUTO_INCREMENT" json:"created_at"`
		UpdatedAt   string         `gorm:"column:updated_at;AUTO_INCREMENT" json:"updated_at"`
		DeletedAt   gorm.DeletedAt `gorm:"AUTO_INCREMENT" json:"deleted_at"`
	}
)

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

func NewUsersResponse(users model.Users) UsersResponse {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		loc = time.UTC
	}
	return UsersResponse{
		Id:          users.Id,
		Name:        users.Name,
		SName:       users.SName,
		Username:    users.Username,
		Password:    users.Password,
		Email:       users.Email,
		Phone:       users.Phone,
		DateOfBirth: users.DateOfBirth,
		Gender:      users.Gender,
		Thumbnail:   users.Thumbnail,
		CreatedAt:   time.Time(users.CreatedAt).In(loc).Format(time.RFC3339),
		UpdatedAt:   time.Time(users.UpdatedAt).In(loc).Format(time.RFC3339),
	}
}
