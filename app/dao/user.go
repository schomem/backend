package dao

import "schomem/app/models"

type UserDaoInterface struct{}

func (dao *UserDaoInterface) UserCreate(user *models.User) error {
	result := DB.Create(user)
	return result.Error
}
