package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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

func InsertProduct(c *gin.Context) {
	db, _ := connectToMariaDB()
	var Product Products
	c.BindJSON(&Product)
	log.Println(Product)
	db.Create(&Product)
	c.JSON(200, gin.H{
		"Message": "Product has been created",
	})
}

func UploadsIMG(c *gin.Context) {
	db, _ := connectToMariaDB()
	id := c.Param("id")
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	var Product Products
	db.First(&Product, "product_id = ?", id)
	c.SaveUploadedFile(file, "public/uploads/"+file.Filename)
	Product.Thumbnail = "/uploads/" + file.Filename
	db.Save(&Product)
	c.JSON(200, gin.H{
		"data": "file uploaded!",
	})
}

func getAllProductsCategory(c *gin.Context) {
	db, _ := connectToMariaDB()
	ptype := c.Param("product_type")
	var Product []Products
	err := db.Where("product_type = ?", ptype).Find(&Product).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"data": "Product not Found",
		})
	}
	c.JSON(200, gin.H{
		"data": Product,
	})
}

func OffsetProducts(c *gin.Context) {
	db, _ := connectToMariaDB()
	page := c.Param("page")
	n, _ := strconv.Atoi(page)
	offset := (n - 1) * 5
	var Product []Products
	db.Offset(offset).Limit(5).Find(&Product)
	c.JSON(200, gin.H{
		"Message": Product,
	})
}
