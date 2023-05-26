package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateNewProduct(t *testing.T) {
	product, err := NewProduct("Luiz Henrique", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.Equal(t, "Luiz Henrique", product.Name)
	assert.Equal(t, 10, product.Price)
}

func Test_ShouldThrowEmptyNameError(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func Test_ShouldThrowEmptyPriceError(t *testing.T) {
	product, err := NewProduct("Luiz Henrique", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func Test_ShouldThrowInvalidPriceError(t *testing.T) {
	product, err := NewProduct("Luiz Henrique", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, ErrInvalidPrice, err)
}

func Test_ShouldCallValidateMethod(t *testing.T) {
	product, err := NewProduct("Luiz Henrique", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
