package implementations

import (
	"github.com/bukhavtsov/restful-app/dao/entities"
	"github.com/bukhavtsov/restful-app/database/connection"
)

type DeveloperDAOImpl struct {
}

func (database DeveloperDAOImpl) Read(id int64) (*entities.Developer, error) {
	db := connection.GetConnection()
	defer db.Close()
	developer := entities.Developer{}
	if err := db.Where("id = ?", id).Find(&developer).Error; err != nil {
		return nil, err
	}
	return &developer, nil
}

func (database DeveloperDAOImpl) ReadAll() ([]*entities.Developer, error) {
	db := connection.GetConnection()
	defer db.Close()
	developers := make([]*entities.Developer, 0)
	if err := db.Find(&developers).Error; err != nil {
		return nil, err
	}
	return developers, nil
}

func (database DeveloperDAOImpl) Create(developer *entities.Developer) (int64, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Create(developer).Error; err != nil {
		return -1, err
	}
	return developer.Id, nil
}
func (database DeveloperDAOImpl) Update(developer *entities.Developer) error {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", developer.Id).Error; err != nil {
		return err
	}
	if err := db.Update(developer).Error; err != nil {
		return err
	}
	return nil
}
func (database DeveloperDAOImpl) Delete(id int64) error {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", id).Delete(entities.Developer{}).Error; err != nil {
		return err
	}
	return nil
}
