package main

import (
	"flag"
	"strconv"

	"schomem/app/consts"
	"schomem/app/controllers"
	"schomem/app/dao"
	"schomem/app/router"
	"schomem/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
	Config     *utils.Config
)

func init() {
	flag.StringVar(&configPath, "path", utils.FullPath("./config.ini"), "配置文件目录")
	flag.Parse()

	consts.CONFIG = utils.ParseConfig(configPath)

	logrus.SetLevel(logrus.TraceLevel)
}

func main() {

	r := gin.Default()

	// 初始化路由路径
	router.InitRouter(r)

	// 初始化数据库
	dao.InitDb()

	// 初始化验证器
	controllers.InitVali()

	// r.SetTrustedProxies([]string{""})

	if consts.CONFIG.SSL.Use {
		r.RunTLS(":"+strconv.Itoa(consts.CONFIG.General.Port), consts.CONFIG.SSL.Pem, consts.CONFIG.SSL.Key)
	} else {
		r.Run(":" + strconv.Itoa(consts.CONFIG.General.Port))
	}
}
