package dao

import (
	"schomem/app/consts"
	"schomem/app/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	UserDao UserDaoInterface
)

func InitDb() {
	var err error

	dsn := consts.CONFIG.Database.User + ":" + consts.CONFIG.Database.Password + "@tcp(" + consts.CONFIG.Database.Host + ":" + strconv.Itoa(consts.CONFIG.Database.Port) + ")/" + consts.CONFIG.Database.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitTables()
}

func InitTables() {
	DB.AutoMigrate(&models.User{})
}
