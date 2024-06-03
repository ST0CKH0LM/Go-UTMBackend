package usecase

import (
	"log"

	"gitlab.com/Std217/test/model"
	"gitlab.com/Std217/test/repositories"
	"gitlab.com/Std217/test/serializers"
)

type ProductUsecase struct {
	productRepo repositories.ProductRepository
}

func NewProductUsecase(repo repositories.ProductRepository) *ProductUsecase {
	return &ProductUsecase{productRepo: repo}
}

func (uc *ProductUsecase) GetAllProducts() ([]serializers.ProductResponse, error) {
	return uc.productRepo.GetAll()
}

func (uc *ProductUsecase) GetAllCatagory() ([]model.Products, error) {
	return uc.productRepo.GetAllCat()
}

func (uc *ProductUsecase) GetProductsDetail(id uint) (*model.Products, error) {
	return uc.productRepo.GetProductsDetail(id)
}

func (uc *ProductUsecase) GetOffSetProducts(nPage, limit int) (int64, []model.Products, error) {
	return uc.productRepo.GetOffSetProducts(nPage, limit)
}

func (uc *ProductUsecase) GetTotalPage() (float64, error) {
	return uc.productRepo.GetTotalPage()
}

func (uc *ProductUsecase) SearchProduct(nOffset int, limit int, sProduct model.Products) ([]model.Products, error) {
	aProduct, err := uc.productRepo.SearchProduct(nOffset, limit, sProduct)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return aProduct, nil
}

func (uc *ProductUsecase) GetAllProductsCategory(sCatagory model.Products) ([]model.Products, error) {
	aProduct, err := uc.productRepo.GetAllProductsCategory(sCatagory)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return aProduct, nil
}

func (uc *ProductUsecase) InsertProduct(product *model.Products) error {
	if err := uc.productRepo.InsertProduct(product); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// func (uc *ProductUsecase) UploadsIMG(id string, filepath string, product *model.Products) error {
// 	if err := uc.productRepo.UploadsIMG(id, filepath, product); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (uc *ProductUsecase) FindProductByID(id string, product *model.Products) error {
	return uc.productRepo.FindProductByID(id, product)
}

func (uc *ProductUsecase) UpdateProduct(product *model.Products) error {
	return uc.productRepo.UpdateProduct(product)
}
