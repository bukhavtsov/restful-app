package daos

import (
	"github.com/bukhavtsov/restful-app/database/connection"
	"github.com/bukhavtsov/restful-app/models"
)

type CustomerDAO struct {
}

func (CustomerDAO) Read(id int64) (*models.Customer, error) {
	db := connection.GetConnection()
	defer db.Close()
	customer := models.Customer{}
	if err := db.Where("id = ?", id).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (CustomerDAO) ReadAll() ([]*models.Customer, error) {
	db := connection.GetConnection()
	defer db.Close()
	customers := make([]*models.Customer, 0)
	if err := db.Find(&customers).Error; err != nil {
		return []*models.Customer{}, err
	}
	return customers, nil
}

func (CustomerDAO) Create(customer *models.Customer) (int64, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Create(customer).Error; err != nil {
		return -1, err
	}
	return customer.Id, nil
}

func (CustomerDAO) Update(customer *models.Customer) (*models.Customer, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (CustomerDAO) Delete(id int64) error {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", id).Delete(&models.Customer{}).Error; err != nil {
		return err
	}
	return nil
}
