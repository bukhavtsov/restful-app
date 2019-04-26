package models

type Customer struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Money    int64  `gorm:"column:money" json:"money"`
	Discount int64  `gorm:"column:discount" json:"discount"`
	State    string `gorm:"column:state" json:"state"`
}
