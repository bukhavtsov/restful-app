package daos

import (
	"github.com/bukhavtsov/restful-app/database/connection"
	"github.com/bukhavtsov/restful-app/models"
)

type DeveloperDAO struct {
}

func (database DeveloperDAO) Read(id int64) (*models.Developer, error) {
	db := connection.GetConnection()
	defer db.Close()
	developer := models.Developer{}
	if err := db.Where("id = ?", id).Find(&developer).Error; err != nil {
		return nil, err
	}
	return &developer, nil
}

func (database DeveloperDAO) ReadAll() ([]*models.Developer, error) {
	db := connection.GetConnection()
	defer db.Close()
	developers := make([]*models.Developer, 0)
	if err := db.Find(&developers).Error; err != nil {
		return []*models.Developer{}, err
	}
	return developers, nil
}

func (database DeveloperDAO) Create(developer *models.Developer) (int64, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Create(developer).Error; err != nil {
		return -1, err
	}
	return developer.Id, nil
}

func (database DeveloperDAO) Update(developer *models.Developer) (*models.Developer, error) {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Save(&developer).Error; err != nil {
		return nil, err
	}
	return developer, nil
}

func (database DeveloperDAO) Delete(id int64) error {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", id).Delete(&models.Developer{}).Error; err != nil {
		return err
	}
	return nil
}
