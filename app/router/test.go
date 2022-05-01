package router

import (
	"schomem/app/consts"
	"schomem/app/controllers"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	r.GET("/hi", func(c *gin.Context) {
		controllers.Result.Ok(c, "ok")
	})

	r.GET("/allowRegister", func(c *gin.Context) {
		controllers.Result.Ok(c, consts.CONFIG.Service.Register)
	})
}
