package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Marcos Augusto", "teste@teste.com", "1234567")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Marcos Augusto", user.Name)
	assert.Equal(t, "teste@teste.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("Marcos Augusto", "teste@teste.com", "1234567")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("1234567"))
	assert.False(t, user.ValidatePassword("4125364"))
	assert.NotEqual(t, user.Password, "12345678")
}
