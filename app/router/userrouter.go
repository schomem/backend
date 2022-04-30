package router

import (
	"schomem/app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	r.Group("/user")
	{
		r.POST("/register", controllers.UserContro.SaveUser)
	}
}
