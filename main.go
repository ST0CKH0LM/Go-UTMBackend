package main

import (
	"fmt"
	"net/http"
	"os"

	"gitlab.com/Std217/test/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	_ *gorm.DB
)

func connectToMariaDB() (*gorm.DB, error) {
	if err := config.EnvLoad(); err != nil {
		println(err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")
	// DB_PASS
	// DB_USER
	// DB_URL
	// DB_NAME
	Path := dbUser + ":" + dbPass + "@tcp(" + dbURL + ")/" + dbName
	// Path := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbURL, dbName)
	println(Path)
	// db, err := sql.Open("mysql", Path)
	db, err := gorm.Open(mysql.Open(Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MariaDB!")
	return db, nil
}

func main() {
	db, _ := connectToMariaDB()
	sqlDB, _ := db.DB()
	sqlDB.Close()
	CreateTable()
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLGlob("templates/*")
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	})
	r.GET("/", index)
	r.GET("/ping", testPing)
	r.GET("/user", getAllUser)
	r.GET("/user/:user_id", getUserById)
	r.GET("/register", InsertUser)
	r.GET("/products", getAllProducts)
	r.GET("/products/:product_id", getProductsDetail)
	r.POST("/uploads", UploadsIMG)
	r.StaticFS("/uploads", http.Dir("public/uploads"))
	r.Run()
}

func CreateTable() (c *gin.Context) {
	db, _ := connectToMariaDB()
	err := db.AutoMigrate(&Products{}, &Persons{})
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
	return
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "test",
	})
}
