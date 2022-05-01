package controllers

import "github.com/gin-gonic/gin"

type Res struct{}

func (result *Res) Ok(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func (result *Res) ParamsErr(c *gin.Context, msg string) {
	c.JSON(400, gin.H{
		"code": 400,
		"msg":  msg,
		"data": gin.H{},
	})
}

func (result *Res) AuthErr(c *gin.Context, msg string) {
	c.JSON(401, gin.H{
		"code": 401,
		"msg":  msg,
		"data": gin.H{},
	})
}

func (result *Res) ServerErr(c *gin.Context, msg string) {
	c.JSON(500, gin.H{
		"code": 500,
		"msg":  msg,
		"data": gin.H{},
	})
}
