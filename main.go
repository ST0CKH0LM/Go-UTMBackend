package main

import (
	"fmt"
	"log"
	"os"

	// "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"gitlab.com/Std217/test/db"
	"gitlab.com/Std217/test/handler"
	"gitlab.com/Std217/test/model"
	"gitlab.com/Std217/test/repositories"
	"gitlab.com/Std217/test/usecase"
)

func main() {
	checkTable()
	database, _ := db.ConnectToMariaDB()
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.LoadHTMLGlob("templates/*")
	r.Use(CORSMiddleware())

	// r.GET("/ping", testPing)
	// r.GET("/user", getAllUser)
	// r.GET("/user/:id", getUserById)
	// r.POST("/register", InsertUser)
	// r.POST("/login", UserLogin)
	// r.POST("/user/:id/img", addProfileImage)
	// r.StaticFS("/uploads", http.Dir("public/uploads"))
	productAPI := r.Group("")
	{
		productRepo := repositories.NewProductRepository(database)
		productUsecase := usecase.NewProductUsecase(*productRepo)
		productHandler := handler.NewProductHandler(*productUsecase)
		productAPI.GET("/search", productHandler.SearchProduct)
		productAPI.GET("/products", productHandler.GetAllProducts)
		productAPI.GET("/producttype", productHandler.GetAllProductsCategory)
		productAPI.GET("/alltype", productHandler.GetAllCatagory)
		productAPI.GET("/products/:id", productHandler.GetProductsDetail)
		productAPI.GET("/offsetproduct", productHandler.GetOffSetProducts)
		// productAPI.GET("/testCatagory", SelectCatagory)
		productAPI.POST("/products/:id/uploads", productHandler.UploadsIMG)
		productAPI.POST("/products/insert", productHandler.InsertProduct)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Println("err : ", err)
	}
	r.Run(":" + os.Getenv("PORT"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Creadentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func checkTable() (c *gin.Context) {
	db, _ := db.ConnectToMariaDB()
	err := db.AutoMigrate(&model.Products{})
	if err != nil {
		fmt.Println(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
	return
}

// func index(c *gin.Context) {
// 	c.HTML(http.StatusOK, "index.html", gin.H{
// 		"title": "test",
// 		"image": "uploads/Ch.jpg",
// 	})
// }
