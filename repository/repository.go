package repository

import (
	"fmt"
	"github.com/babay15/product-app/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Repository interface {
	SaveProduct(product entity.Product)
	UpdateProduct(product entity.Product)
	DeleteProduct(product entity.Product)
	FindAllProduct() []entity.Product
	SaveBrand(brand entity.Brand)
	UpdateBrand(brand entity.Brand)
	DeleteBrand(brand entity.Brand)
	FindAllBrand() []entity.Brand
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewRepository() Repository{
	dsn := CreateDSN()
	db, err := gorm.Open(postgres.Open(dsn))
	if err!=nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.Brand{})
	return &database {
		connection: db,
	}
}

func (d *database) SaveProduct(product entity.Product) {
	d.connection.Create(&product)
}

func (d *database) UpdateProduct(product entity.Product) {
	d.connection.Save(&product)
}

func (d *database) DeleteProduct(product entity.Product) {
	d.connection.Delete(&product)
}

func (d *database) FindAllProduct() []entity.Product {
	var products []entity.Product
	d.connection.Find(&products)
	return products
}

func (d *database) SaveBrand(brand entity.Brand) {
	d.connection.Create(&brand)
}

func (d *database) UpdateBrand(brand entity.Brand) {
	d.connection.Save(&brand)
}

func (d *database) DeleteBrand(brand entity.Brand) {
	d.connection.Delete(&brand)
}

func (d *database) FindAllBrand() []entity.Brand {
	var brands []entity.Brand
	d.connection.Find(&brands)
	return brands
}

func (d *database) CloseDB() {
	err := d.connection.Clauses()
	if err!=nil{
		panic("Failed to close the database")
	}
}

func CreateDSN() string {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "home_business"
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
}
