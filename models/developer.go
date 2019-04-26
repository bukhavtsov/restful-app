package models

type Developer struct {
	Id           int64  `gorm:"column:id" json:"id"`
	Name         string `gorm:"column:name" json:"name"`
	Age          int64  `gorm:"column:age" json:"age"`
	PrimarySkill string `gorm:"column:primary_skill" json:"primary_skill"`
}
