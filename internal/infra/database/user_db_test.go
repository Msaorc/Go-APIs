package database

import (
	"testing"

	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory.db"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Marcos Augusto", "marcos@email.com", "1234567")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFinded entity.User
	err = db.First(&userFinded, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFinded.ID)
	assert.Equal(t, user.Name, userFinded.Name)
	assert.Equal(t, user.Email, userFinded.Email)
	assert.NotNil(t, userFinded.Password)
}
