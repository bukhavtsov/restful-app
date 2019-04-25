package entities

type Developer struct {
	Id           int64  `gorm:"AUTO_INCREMENT";"column:beast_id"`
	Name         string `gorm:"not null"`
	Age          int64  `gorm:"not null"`
	PrimarySkill string `gorm:"not null"`
}
