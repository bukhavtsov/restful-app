package data

import (
	"github.com/bukhavtsov/restful-app/pkg/database/connection"
	"github.com/bukhavtsov/restful-app/pkg/models"
)

type userDAO struct{}

func NewUserDAO() *userDAO {
	return &userDAO{}
}

func (dao *userDAO) Get(login, password string) (*models.User, error) {
	db := connection.GetConnection()
	defer db.Close()
	user := models.User{}
	if err := db.Where("login = ? AND password = ?", login, password).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAO) GetById(id int64) (*models.User, error) {
	db := connection.GetConnection()
	defer db.Close()
	user := models.User{}
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *userDAO) Create(user *models.User) (int64, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Create(user).Error; err != nil {
		return -1, err
	}
	return user.Id, nil
}

func (dao *userDAO) Update(user *models.User, refreshToken string) (*models.User, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Model(&user).Where("id = ?", user.Id).Update("refresh_token", refreshToken).Error; err != nil {
		return nil, err
	}
	return user, nil
}
