package data

import (
	"github.com/jinzhu/gorm"
)

type Customer struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Money    int64  `gorm:"column:money" json:"money"`
	Discount int64  `gorm:"column:discount" json:"discount"`
	State    string `gorm:"column:state" json:"state"`
}

type customerData struct {
	db *gorm.DB
}

func NewCustomerData(db *gorm.DB) *customerData {
	return &customerData{db}
}

func (d *customerData) Read(id int64) (*Customer, error) {
	customer := Customer{}
	if err := d.db.Where("id = ?", id).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (d *customerData) ReadAll() ([]*Customer, error) {
	customers := make([]*Customer, 0)
	if err := d.db.Find(&customers).Error; err != nil {
		return []*Customer{}, err
	}
	return customers, nil
}

func (d *customerData) Create(customer *Customer) (int64, error) {
	if err := d.db.Create(customer).Error; err != nil {
		return -1, err
	}
	return customer.Id, nil
}

func (d *customerData) Update(customer *Customer) (*Customer, error) {
	if err := d.db.Save(&customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}

func (d *customerData) Delete(id int64) error {
	if err := d.db.Where("id = ?", id).Delete(&Customer{}).Error; err != nil {
		return err
	}
	return nil
}
