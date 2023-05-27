package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/MogLuiz/key-user-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Test_E2E_ShouldCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 100.00, product.Price)
}

func Test_E2E_ShouldFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)

	err = productDB.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func Test_E2E_ShouldUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Product 2"
	product.Price = 200.00
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", productFound.Name)
	assert.Equal(t, 200.00, productFound.Price)
}

func Test_E2E_ShouldDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 100.00)
	assert.Nil(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}

func Test_E2E_ShouldFindAllProductsWithPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	productsPage1, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(productsPage1))
	assert.Equal(t, "Product 1", productsPage1[0].Name)
	assert.Equal(t, "Product 10", productsPage1[9].Name)

	productsPage2, err := productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(productsPage2))
	assert.Equal(t, "Product 11", productsPage2[0].Name)
	assert.Equal(t, "Product 20", productsPage2[9].Name)

	productsPage3, err := productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 4, len(productsPage3))
	assert.Equal(t, "Product 21", productsPage3[0].Name)
	assert.Equal(t, "Product 24", productsPage3[3].Name)
}

func Test_E2E_ShouldFindAllProductsWithoutPagination(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})

	for i := 1; i <= 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(0, 0, "asc")
	assert.NoError(t, err)
	assert.Equal(t, 24, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 24", products[23].Name)
}
