package entities

type Customer struct {
	Id       int64  `gorm:"column:id; not null" json:"id"`
	Name     string `gorm:"column:name; not null" json:"name"`
	Money    int64  `gorm:"column:money; not null" json:"money"`
	Discount int64  `gorm:"column:discount; not null" json:"discount"`
	State    string `gorm:"column:state; not null" json:"state"`
}
