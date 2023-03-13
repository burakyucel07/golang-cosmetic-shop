package repositories

import (
	"burakyucel/test/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type ProductRepository interface {
	Save(product entities.Product)
	Update(product entities.Product)
	Delete(product entities.Product)
	FindAll() []entities.Product
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() ProductRepository {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect database.")
	}

	db.AutoMigrate(&entities.Product{})

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()

	if err != nil {
		panic("Failed to close database.")
	}
}

func (db *database) Save(product entities.Product) {
	db.connection.Create(&product)
}

func (db *database) Update(product entities.Product) {
	db.connection.Save(&product)
}

func (db *database) Delete(product entities.Product) {
	db.connection.Delete(&product)
}

func (db *database) FindAll() []entities.Product {
	var products []entities.Product
	db.connection.Find(&products)

	return products
}
