package interfaces

import "github.com/bukhavtsov/restful-app/dao/entities"

type CustomerDAO interface {
	Create(customer *entities.Customer) (int64, error)
	Read(id int64) (*entities.Customer, error)
	ReadAll() ([]*entities.Customer, error)
	Update(customer *entities.Customer) (*entities.Customer, error)
	Delete(id int64) error
}
