package dao

import (
	"schomem/app/models"
	"time"
)

type UserDaoInterface struct{}

func (dao *UserDaoInterface) Create(user *models.User) error {
	user.PasswordUpdatedAt = time.Now()
	result := DB.Create(user)
	return result.Error
}

func (dao *UserDaoInterface) DeleteOneById(id int) error {
	user := dao.FindOneById(id)
	return dao.Delete(&user)
}

func (dao *UserDaoInterface) Delete(user *models.User) error {
	result := DB.Delete(user)
	return result.Error
}

func (dao *UserDaoInterface) FindOneById(id int) models.User {
	var user models.User
	DB.First(&user, id)
	return user
}

func (dao *UserDaoInterface) FindOneByCustomId(custom_id string) models.User {
	var user models.User
	DB.Where("custom_id = ?", custom_id).First(&user)
	return user
}

func (dao *UserDaoInterface) FindOneByEmail(email string) models.User {
	var user models.User
	DB.Where("email = ?", email).First(&user)
	return user
}
