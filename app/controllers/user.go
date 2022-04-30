package controllers

import (
	"schomem/app/models"
	"schomem/app/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct{}

func (t *UserController) FindAllUser(c *gin.Context) {}

func (t *UserController) SaveUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user = models.User{Name: "test", Email: "test@test.test", Password: "test"}
	if err := services.UserService.SaveUser(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success"})
	}
}
