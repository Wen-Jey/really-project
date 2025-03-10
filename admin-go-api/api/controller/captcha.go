//验证码控制层

package controller

import (
	"admin-go-api/api/service"
	"admin-go-api/common/result"
	"github.com/gin-gonic/gin"
)

// @summary 验证码接口

//  @Produce json

// @Description 验证码接口

// @Success 200 {object} result.Result

// Captcha @router /api/captcha [get]
func Captcha(c *gin.Context) {
	id, base64Image := service.CaptMake()
	result.Success(c, map[string]interface{}{"idKey": id, "Image": base64Image})
}
