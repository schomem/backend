package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func InitVali() {
	var ok bool
	v, ok = binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("validator engine is not validator.Validate")
	} else {
		// 自定义验证规则
	}
}

func ParseQueryRequest(c *gin.Context, data interface{}) error {
	if err := c.ShouldBindQuery(data); err != nil {
		Result.ParamsErr(c, err.Error())
		return err
	}
	return nil
}

func ParseJSONRequest(c *gin.Context, data interface{}) error {
	if err := c.ShouldBind(data); err != nil {
		Result.ParamsErr(c, err.Error())
		return err
	}
	return nil
}
