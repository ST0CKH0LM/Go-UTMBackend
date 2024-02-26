package main

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Persons struct {
	// gorm.Model
	User_ID     int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string `gorm:"not null"`
	SName       string
	Username    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Email       string
	DateOfBirth string
	Gender      string
	Thumbnail   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func getAllUser(c *gin.Context) {
	db, _ := connectToMariaDB()
	var Users []Persons
	db.Find(&Users)
	c.JSON(200, gin.H{
		"data": Users,
	})
}

func getUserById(c *gin.Context) {
	db, _ := connectToMariaDB()
	id := c.Param("user_id")
	var Person []Persons
	err := db.First(&Person, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"data": "Error! User not found",
		})
	} else {
		c.JSON(200, gin.H{
			"data": Person,
		})
	}
}

func InsertUser(c *gin.Context) {
	db, _ := connectToMariaDB()
	SecretPass := "testPassword"
	hash, _ := HashPassword(SecretPass)
	newUser := Persons{
		Name:        "testName",
		SName:       "testSName",
		Username:    "testUsername",
		Password:    hash,
		Email:       "testEmail@gmail.com",
		DateOfBirth: "2000/02/08",
		Gender:      "Male",
	}
	db.Create(&newUser)
	c.JSON(200, gin.H{
		"data": "User Create",
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
