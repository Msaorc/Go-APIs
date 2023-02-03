package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateTableProductAndBD() *Product {
	db, err := gorm.Open(sqlite.Open("file:memory.db"))
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(entity.Product{})
	db.AutoMigrate(&entity.Product{})
	return NewProduct(db)
}

func TestCreateNewProduct(t *testing.T) {
	productDB := CreateTableProductAndBD()
	product, err := entity.NewProduct("notebook", 10000.00)
	assert.Nil(t, err)
	err = productDB.Create(product)
	assert.Nil(t, err)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "notebook", product.Name)
	assert.Equal(t, 10000.00, product.Price)
}

func TestFindAllProducts(t *testing.T) {
	productDB := CreateTableProductAndBD()
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.Nil(t, err)
		productDB.Create(product)
	}
	products, err := productDB.FindAll(1, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	productDB := CreateTableProductAndBD()
	product, err := entity.NewProduct("notebook", 16000.00)
	assert.Nil(t, err)
	productDB.Create(product)
	product, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "notebook", product.Name)
	assert.Equal(t, 16000.00, product.Price)
}

func TestUpdateProduct(t *testing.T) {
	productDB := CreateTableProductAndBD()
	product, err := entity.NewProduct("notebook", 16000.00)
	assert.Nil(t, err)
	productDB.Create(product)
	product.Name = "Notebook Dell"
	assert.NotNil(t, product.ID)
	err = productDB.Update(product)
	assert.Nil(t, err)
	product, err = productDB.FindByID(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, "Notebook Dell", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	productDB := CreateTableProductAndBD()
	product, err := entity.NewProduct("notebook", 16000.00)
	assert.Nil(t, err)
	productDB.Create(product)
	err = productDB.Delete(product.ID.String())
	assert.Nil(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
