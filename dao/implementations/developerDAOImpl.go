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
		return []*entities.Developer{}, err
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

func (database DeveloperDAOImpl) Update(developer *entities.Developer) (*entities.Developer, error) {
	db := connection.GetConnection()
	defer db.Close()
	var newDeveloper entities.Developer
	db.First(&newDeveloper)
	newDeveloper.Id = developer.Id
	newDeveloper.PrimarySkill = developer.PrimarySkill
	newDeveloper.Age = developer.Age
	newDeveloper.Name = developer.Name
	if err := db.Save(&newDeveloper).Error; err != nil {
		return nil, err
	}
	return &newDeveloper, nil
}
func (database DeveloperDAOImpl) Delete(id int64) error {
	db := connection.GetConnection()
	defer db.Close()
	if err := db.Where("id = ?", id).Delete(&entities.Developer{}).Error; err != nil {
		return err
	}
	return nil
}
