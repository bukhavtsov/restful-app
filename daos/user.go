package daos

import (
	"github.com/bukhavtsov/restful-app/database/connection"
	"github.com/bukhavtsov/restful-app/models"
)

type UserDAO struct{}

func (UserDAO) Get(login, password string) (*models.User, error) {
	db := connection.GetConnection()
	defer db.Close()
	user := models.User{}
	if err := db.Where("login = ? AND password = ?", login, password).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (UserDAO) Create(user *models.User) (int64, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Create(user).Error; err != nil {
		return -1, err
	}
	return user.Id, nil
}
