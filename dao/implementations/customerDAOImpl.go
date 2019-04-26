package implementations

import (
	"github.com/bukhavtsov/restful-app/dao/entities"
	"github.com/bukhavtsov/restful-app/database/connection"
)

type CustomerDAOImpl struct {
}

func (database CustomerDAOImpl) Read(id int64) (*entities.Customer, error) {
	db := connection.GetConnection()
	defer db.Close()
	customer := entities.Customer{}
	if err := db.Where("id = ?", id).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (database CustomerDAOImpl) ReadAll() ([]*entities.Customer, error) {
	db := connection.GetConnection()
	defer db.Close()
	customers := make([]*entities.Customer, 0)
	if err := db.Find(&customers).Error; err != nil {
		return []*entities.Customer{}, err
	}
	return customers, nil
}

func (database CustomerDAOImpl) Create(customer *entities.Customer) (int64, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Create(customer).Error; err != nil {
		return -1, err
	}
	return customer.Id, nil
}

func (database CustomerDAOImpl) Update(customer *entities.Customer) (*entities.Customer, error) {
	db := connection.GetConnection()
	defer db.Close()
	var newCustomer entities.Customer
	db.First(&newCustomer)
	newCustomer.Id = customer.Id
	newCustomer.Name = customer.Name
	newCustomer.Discount = customer.Discount
	newCustomer.Money = customer.Money
	newCustomer.State = customer.State
	if err := db.Save(&newCustomer).Error; err != nil {
		return nil, err
	}
	return &newCustomer, nil
}
func (database CustomerDAOImpl) Delete(id int64) error {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", id).Delete(&entities.Customer{}).Error; err != nil {
		return err
	}
	return nil
}
