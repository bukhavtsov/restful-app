package entities

type Developer struct {
	Id           int64  `gorm:"column:id; not null" json:"id"`
	Name         string `gorm:"column:name; not null" json:"name"`
	Age          int64  `gorm:"column:age; not null" json:"age"`
	PrimarySkill string `gorm:"column:primary_skill; not null" json:"primary_skill"`
}
