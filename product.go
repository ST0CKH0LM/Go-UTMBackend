package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Products struct {
	// gorm.Model
	Product_id    uint   `gorm:"primaryKey;AUTO_INCREMENT"`
	Product_name  string `gorm:"not null"`
	Product_type  string `gorm:"not null"`
	Product_price int    `gorm:"not null"`
	Product_count int    `gorm:"not null"`
	Product_brand string ``
	Thumbnail     string ``
	Rate          int    ``
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func getAllProducts(c *gin.Context) {
	db, _ := connectToMariaDB()
	var Product []Products
	db.Find(&Product)
	c.JSON(200, gin.H{
		"list": Product,
	})
}

func getProductsDetail(c *gin.Context) {
	id := c.Param("product_id")
	db, _ := connectToMariaDB()
	fmt.Println(id)
	var Product Products
	err := db.First(&Product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"data": "Product not found",
		})
	}
	c.JSON(200, gin.H{
		"data": Product,
	})
}

func UploadsIMG(c *gin.Context) {
	db, _ := connectToMariaDB()
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	c.SaveUploadedFile(file, "public/uploads/"+file.Filename)
	newUploads := Products{
		Product_name:  "test",
		Product_type:  "test",
		Product_price: 300,
		Product_count: 2000,
		Product_brand: "test",
		Rate:          5,
		CreatedAt:     time.Now(),
		Thumbnail:     "/uploads/" + file.Filename,
	}
	db.Create(&newUploads)
	c.JSON(200, gin.H{
		"data": "file uploaded!",
	})
}
