package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		UserRouter(v1)
		TestRouter(v1)
	}
}
