//岗位控制层 针对service进行调用

package controller

import (
	"admin-go-api/api/entity"
	"admin-go-api/api/service"
	"github.com/gin-gonic/gin"
)

var sysPost entity.SysPost

// 新增岗位
// @Summary 新增岗位接口
// @Produce json
// @Description 新增岗位接口
// @Param data body entity.SysPost true "data"
// @Success 200 {object} result.Result
// @Router /api/post/add [post]
func CreateSysPost(c *gin.Context) {
	_ = c.BindJSON(&sysPost)
	service.SysPostService().CreateSysPost(c, sysPost)
}
