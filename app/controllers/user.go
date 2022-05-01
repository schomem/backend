package controllers

import (
	"schomem/app/models"
	"schomem/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (t *UserController) Login(c *gin.Context) {
	var user models.UserLogin
	if err := ParseQueryRequest(c, &user); err != nil {
		return
	}
	services.UserService.Login(&user)
	Result.Ok(c, "ok")
}

func (t *UserController) Register(c *gin.Context) {
	var user models.User
	if err := ParseJSONRequest(c, &user); err != nil {
		return
	}
	if err := services.UserService.Register(&user); err != nil {
		Result.ServerErr(c, err.Error())
	} else {
		Result.Ok(c, "ok")
	}
}
