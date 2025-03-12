//岗位服务层

package service

import (
	"admin-go-api/api/dao"
	"admin-go-api/api/entity"
	"admin-go-api/common/result"
	"github.com/gin-gonic/gin"
)

type ISysPostService interface {
	CreateSysPost(c *gin.Context, sysPost entity.SysPost)
}

type SysPostServiceImpl struct{}

//新增岗位

func (s SysPostServiceImpl) CreateSysPost(c *gin.Context, sysPost entity.SysPost) {
	bool := dao.CreateSysPost(sysPost)
	if !bool {
		result.Failed(c, int(result.ApiCode.POSTALREADYEXISTS), result.ApiCode.GetMessage(result.ApiCode.POSTALREADYEXISTS))
		return
	}
	result.Success(c, true)
}

var sysPostService = SysPostServiceImpl{}

func SysPostService() ISysPostService {
	return &sysPostService
}
