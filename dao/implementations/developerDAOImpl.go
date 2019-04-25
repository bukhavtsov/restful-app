package implementations

import (
	. "../../database/connection"
	. "../entities"
	"github.com/bukhavtsov/restful-app/dao/entities"
)

type DeveloperDAOImpl struct {
}

func (database DeveloperDAOImpl) Read(id int64) (*Developer, error) {
	db := GetConnection()
	defer db.Close()
	developer := entities.Developer{}
	if err := db.Where("id = ?", id).Find(&developer).Error; err != nil {
		return nil, err
	}
	return &developer, nil
}

func (database DeveloperDAOImpl) ReadAll() ([]*Developer, error) {
	db := GetConnection()
	defer db.Close()
	developers := make([]*Developer, 0)
	if err := db.Find(&developers).Error; err != nil {
		return nil, err
	}
	return developers, nil
}

func (database DeveloperDAOImpl) Create(developer *Developer) (int64, error) {
	db := GetConnection()
	defer db.Close()
	if err := db.Create(developer).Error; err != nil {
		return -1, err
	}
	return developer.Id, nil
}
func (database DeveloperDAOImpl) Update(developer *Developer) error {
	db := GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", developer.Id).Update(developer).Error; err != nil {
		return err
	}
	return nil
}
func (database DeveloperDAOImpl) Delete(id int64) error {
	db := GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", id).Delete(Developer{}).Error; err != nil {
		return err
	}
	return nil
}
