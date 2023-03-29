package database

import (
	"testing"

	"github.com/Msaorc/Go-APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateTableUserAndDB() *User {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(entity.User{})
	db.AutoMigrate(&entity.User{})
	return NewUser(db)
}

func TestCreateUser(t *testing.T) {
	userDB := CreateTableUserAndDB()
	user, _ := entity.NewUser("Marcos Augusto", "marcos@email.com", "1234567")
	err := userDB.Create(user)
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Marcos Augusto", user.Name)
	assert.Equal(t, "marcos@email.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestFindByEmail(t *testing.T) {
	userDB := CreateTableUserAndDB()
	user, _ := entity.NewUser("Teste User Find Email", "findemail@email.com", "1234567")
	err := userDB.Create(user)
	assert.Nil(t, err)

	userFind, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFind.ID)
	assert.Equal(t, user.Name, userFind.Name)
	assert.Equal(t, user.Email, userFind.Email)
	assert.NotNil(t, userFind.Password)
}
