package repositories

import (
	"log"
	"math"

	"gitlab.com/Std217/test/model"
	"gitlab.com/Std217/test/serializers"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (repo *ProductRepository) GetAll() ([]serializers.ProductResponse, error) {
	var products []model.Products
	if err := repo.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	var serializedProducts []serializers.ProductResponse
	for _, product := range products {
		serializedProducts = append(serializedProducts, serializers.NewProductResponse(product))
	}
	return serializedProducts, nil
}

func (repo *ProductRepository) SearchProduct(nOffset int, limit int, sProduct model.Products) ([]model.Products, error) {
	var aProduct []model.Products
	if err := repo.DB.Where("Product_name LIKE ?", sProduct.Product_name+"%").Offset(nOffset).Limit(limit).Find(&aProduct).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(sProduct.Product_name)

	return aProduct, nil
}

func (repo *ProductRepository) GetAllProductsCategory(sCatagory model.Products) ([]model.Products, error) {
	var aProduct []model.Products
	if err := repo.DB.Where("Product_type = ?", sCatagory.Product_type).Find(&aProduct).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(sCatagory.Product_type)

	return aProduct, nil
}

func (repo *ProductRepository) GetAllCat() ([]model.Products, error) {
	var aCatagory []model.Products
	if err := repo.DB.Table("products").Select("product_type").Find(&aCatagory).Error; err != nil {
		return nil, err
	}

	return aCatagory, nil
}

func (repo *ProductRepository) GetProductsDetail(id uint) (*model.Products, error) {
	var sProduct model.Products
	if err := repo.DB.First(&sProduct, id).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return &sProduct, nil
}

func (repo *ProductRepository) GetOffSetProducts(nPage, limit int) (int64, []model.Products, error) {
	var aProduct []model.Products
	totalItem := repo.DB.Offset(nPage).Limit(limit).Find(&aProduct).RowsAffected
	return totalItem, aProduct, nil
}

func (repo *ProductRepository) GetTotalPage() (float64, error) {
	var aProduct []model.Products
	set := math.Ceil(float64(repo.DB.Find(&aProduct).RowsAffected) / 10)
	return set, nil
}

func (repo *ProductRepository) InsertProduct(product *model.Products) error {
	if err := repo.DB.Create(&product).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// func (repo *ProductRepository) UploadsIMG(id string, fileName string, Product *model.Products) error {
// 	if err := repo.DB.First(&Product, "id = ?", id).Error; err != nil {
// 		return err
// 	}
// 	Product.Product_gallery = fileName
// 	Product.UpdatedAt = model.JSONTime(time.Now())
// 	if err := repo.DB.Save(&Product).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func (repo *ProductRepository) FindProductByID(id string, product *model.Products) error {
	return repo.DB.First(product, id).Error
}

func (repo *ProductRepository) UpdateProduct(product *model.Products) error {
	return repo.DB.Save(product).Error
}
