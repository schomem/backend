package models

import "time"

type User struct {
	Id                int    `gorm:"primary_key"`
	Custom_Id         string `gorm:"type:varchar(20);unique_index"`
	Gender            int    `gorm:"type:tinyint(1)" json:"gender" binding:"required"` // 0 为男，1 为女，2 为保密
	Name              string `gorm:"type:varchar(100);not null" json:"name" binding:"required"`
	Email             string `gorm:"type:varchar(100);unique_index" json:"email" binding:"required,email"`
	Salt              string `gorm:"type:varchar(8);not null"`
	Password          string `gorm:"type:varchar(100)" json:"password" binding:"required"`
	IsBan             bool   `gorm:"type:tinyint(1)"`
	BanReason         string `gorm:"type:varchar(255)"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	PasswordUpdatedAt time.Time `gorm:"type:datetime"`
}

type UserLogin struct {
	Email    string `form:"email" binding:"email"`
	Name     string `form:"name"`
	Password string `form:"password" binding:"required"`
}
