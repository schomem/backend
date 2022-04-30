package services

import (
	"schomem/app/dao"
	"schomem/app/models"
)

type UserMgr struct{}

type UserServices interface {
	SaveUser(user *models.User) error
}

func (mgr *UserMgr) SaveUser(user *models.User) error {
	return dao.UserDao.UserCreate(user)
}
