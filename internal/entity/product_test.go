package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("NoteBook", 15000)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "NoteBook", p.Name)
	assert.Equal(t, 15000.00, p.Price)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("NoteBook", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("NoteBook", -20)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductWhenNameIsInvalid(t *testing.T) {
	p, err := NewProduct("", 15000)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("NoteBook", 15000)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
