package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Persons struct {
	// gorm.Model
	ID          int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Name        string `gorm:"not null"`
	SName       string
	Username    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Phone       string `gorm:"not null"`
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
	id := c.Param("id")
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
	var Person Persons
	c.BindJSON(&Person)
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(Person.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	Person.Password = string(hashpassword)
	var checkUser Persons
	var checkEmail Persons
	db.First(&checkUser, "username = ?", Person.Username)
	db.First(&checkEmail, "email = ?", Person.Email)
	log.Println(checkUser.Username)
	log.Println(checkUser.Email)
	if checkUser.Username == Person.Username || checkEmail.Email == Person.Email {
		c.JSON(400, gin.H{
			"Message": "Username Already in Use",
		})
	} else {
		db.Create(&Person)
		c.JSON(200, gin.H{
			"Message": "User has been Created",
		})
	}
}

func UserLogin(c *gin.Context) {
	db, _ := connectToMariaDB()
	var Person Persons
	c.BindJSON(&Person)
	var checkUser Persons
	db.First(&checkUser, "username = ?", Person.Username)
	checkHash := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(Person.Password))
	if checkUser.Username == Person.Username && checkHash == nil {
		c.JSON(200, gin.H{
			"Message": "Login Success, Welcome " + Person.Username,
		})
	} else {
		c.JSON(400, gin.H{
			"Message": "Login Failed",
		})
	}
}

func addProfileImage(c *gin.Context) {
	db, _ := connectToMariaDB()
	id := c.Param("id")
	file, _ := c.FormFile("file")
	var Person Persons
	db.First(&Person, "id = ?", id)
	log.Println(Person.Thumbnail)
	if Person.Thumbnail != "/profileimg/"+file.Filename {
		rem, err := os.Stat("public" + Person.Thumbnail)
		log.Println(rem)
		if err == nil {
			os.Chmod("public/profileimg", 0777)
			err = os.Remove("pubilc" + Person.Thumbnail)
			if err != nil {
				log.Println(err)
			}
		}
	}
	c.SaveUploadedFile(file, "public/profileimg/"+file.Filename)
	imgPath := "/profileimg/" + file.Filename
	Person.Thumbnail = imgPath
	db.Save(&Person)
	c.JSON(200, gin.H{
		"Message": "Upload Complete",
	})
}
