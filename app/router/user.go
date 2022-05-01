package router

import (
	"schomem/app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	account := r.Group("/account")
	{
		account.POST("/register", controllers.UserContro.Register)
		account.POST("/login", controllers.UserContro.Login)
	}
}
