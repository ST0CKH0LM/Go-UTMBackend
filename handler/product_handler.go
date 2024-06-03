package handler

import (
	"encoding/json"
	"log"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/Std217/test/model"
	"gitlab.com/Std217/test/usecase"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: usecase}
}

func (handler *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := handler.productUsecase.GetAllProducts()
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(200, gin.H{
		"data": products,
	})
}

func (handler *ProductHandler) GetAllCatagory(c *gin.Context) {
	aCatagory, err := handler.productUsecase.GetAllCatagory()
	if err != nil {
		log.Println(err)
		return
	}

	c.JSON(200, aCatagory)
}

func (handler *ProductHandler) SearchProduct(c *gin.Context) {
	getsProduct := c.Query("product_name")
	sProduct := model.Products{Product_name: getsProduct}
	Offset := c.Query("offset")
	if Offset == "" {
		c.JSON(400, gin.H{"error": "Empty"})
		return
	}
	nPage, err := strconv.Atoi(Offset)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	limit := 10
	nOffset := (nPage - 1) * limit
	aProduct, err := handler.productUsecase.SearchProduct(nOffset, limit, sProduct)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": aProduct})
}

func (handler *ProductHandler) GetProductsDetail(c *gin.Context) {
	idraw := c.Param("id")
	id, err := strconv.ParseUint(idraw, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sProduct, err := handler.productUsecase.GetProductsDetail(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"message": "Not Found"})
	}
	c.JSON(200, sProduct)
}

func (handler *ProductHandler) GetAllProductsCategory(c *gin.Context) {
	var sCatagory model.Products
	if err := c.BindJSON(&sCatagory); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
	}

	aProduct, err := handler.productUsecase.GetAllProductsCategory(sCatagory)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(200, gin.H{
		"data": aProduct,
	})
}

func (handler *ProductHandler) GetOffSetProducts(c *gin.Context) {
	var Page model.Number
	if err := c.BindJSON(&Page); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	nPage := Page.Offsetnumber
	limit := 10
	nOffset := (nPage - 1) * limit
	totalItem, aProduct, err := handler.productUsecase.GetOffSetProducts(nOffset, limit)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	totalPage, err := handler.productUsecase.GetTotalPage()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"currentPage": 0,
		"totalPage":   totalPage,
		"totalItem":   totalItem,
		"list":        aProduct,
	})
}

func (handler *ProductHandler) InsertProduct(c *gin.Context) {
	var product model.Products
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := handler.productUsecase.InsertProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Product has been Created"})
}

// func (handler *ProductHandler) UploadsIMG(c *gin.Context) { // ใช้สำหรับอัพโหลดภาพ Profile อย่าลืมเปลี่ยน DB
// 	var product model.Products
// 	id := c.Param("id")
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.SaveUploadedFile(file, "public/uploads/"+file.Filename)
// 	filepath := "/uploads/" + file.Filename
// 	if err := handler.productUsecase.UploadsIMG(id, filepath, &product); err != nil {
// 		c.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(200, gin.H{"message": "Upload Complete"})
// }

func (handler *ProductHandler) MultiUploads(c *gin.Context) {
	id := c.Param("id")
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	files := form.File["file"]
	var Urls []string

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, "public/uploads/"+filename); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		Urls = append(Urls, "/public/uploads/"+filename)
	}
	var product model.Products
	if err := handler.productUsecase.FindProductByID(id, &product); err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
	}

	imgJSON, err := json.Marshal(Urls)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	product.Product_gallery = imgJSON

	if err := handler.productUsecase.UpdateProduct(&product); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"data":    "Product " + id + " Upload Complete",
		"message": Urls,
	})
}
