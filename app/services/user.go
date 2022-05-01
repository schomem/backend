package services

import (
	"schomem/app/dao"
	"schomem/app/models"
	"strconv"

	"github.com/sirupsen/logrus"
)

type UserMgr struct{}

type UserServices interface {
	Register(user *models.User) error
	Login(user *models.User) string
}

func (mgr *UserMgr) Register(user *models.User) error {
	if err := dao.UserDao.Create(user); err != nil {
		return err
	}
	user.Custom_Id = "id_" + strconv.Itoa(user.Id)
	return nil
}

func (mgr *UserMgr) Login(user *models.UserLogin) string {
	logrus.Debug(dao.UserDao.FindOneByEmail(user.Email))
	return ""
}
