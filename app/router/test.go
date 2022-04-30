package router

import "github.com/gin-gonic/gin"

func TestRouter(r *gin.RouterGroup) {
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
}
