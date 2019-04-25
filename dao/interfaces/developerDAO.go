package interfaces

import "github.com/bukhavtsov/restful-app/dao/entities"

type DeveloperDAO interface {
	Create(developer *entities.Developer) (int64, error)
	Read(id int64) (*entities.Developer, error)
	ReadAll() ([]*entities.Developer, error)
	Update(developer *entities.Developer) error
	Delete(id int64) error
}
