package models

type User struct {
	Id       int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(100)" json:"password" validate:"required"`
}
