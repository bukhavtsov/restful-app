package models

type User struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Login    string `gorm:"column:login" json:"login"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Role     string `gorm:"column:role" json:"role"`
}
