package interfaces

import "../entities"

type DeveloperDAO interface {
	Create(country *entities.Developer) (int64, error)
	Read(id int64) (*entities.Developer, error)
	ReadAll() ([]*entities.Developer, error)
	Update(country *entities.Developer) error
	Delete(id int64) error
}
