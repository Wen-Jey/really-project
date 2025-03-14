//用户控制层

package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"github.com/gin-gonic/gin"
)

// @Summary 用户登录接口
// @Produce json
// @Description 用户登录接口
// @Param data body entity.LoginDto true "data"
// @Success 200 {object} result.Result
// @Router /api/login [post]
func Login(c *gin.Context) {
	var dto entity.LoginDto
	_ = c.ShouldBind(&dto)
	service.SysAdminService().Login(c, dto)
}
