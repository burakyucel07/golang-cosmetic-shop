package services

import (
	"burakyucel/test/entities"
	"burakyucel/test/repositories"
)

type ProductService interface {
	Store(entities.Product) entities.Product
	Update(product entities.Product)
	Delete(product entities.Product)
	FindAll() []entities.Product
}

type productService struct {
	productRepository repositories.ProductRepository
}

func New(repo repositories.ProductRepository) ProductService {
	return &productService{
		productRepository: repo,
	}
}

func (service *productService) Update(product entities.Product) {
	service.productRepository.Update(product)
}

func (service *productService) Delete(product entities.Product) {
	service.productRepository.Delete(product)
}

func (service *productService) Store(product entities.Product) entities.Product {
	service.productRepository.Save(product)

	return product
}

func (service *productService) FindAll() []entities.Product {
	return service.productRepository.FindAll()
}
