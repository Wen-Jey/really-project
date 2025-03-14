//岗位数据层

package dao

import (
	"admin-go-api/api/entity"
	"admin-go-api/common/util"
	. "admin-go-api/pkg/db"
	"time"
)

// 根据编码查询(主要不能重复)

func GetSysPostByCode(postCode string) (sysPost entity.SysPost) {
	Db.Where("post_code = ?", postCode).First(&sysPost)
	return sysPost
}

// 根据名称查询

func GetSysPostByName(postName string) (sysPost entity.SysPost) {
	Db.Where("post_name = ?", postName).First(&sysPost)
	return sysPost
}

// 新增岗位

func CreateSysPost(sysPost entity.SysPost) bool {
	sysPostByCode := GetSysPostByCode(sysPost.PostCode)
	if sysPostByCode.ID > 0 {
		return false
	}
	sysPostByName := GetSysPostByName(sysPost.PostName)
	if sysPostByName.ID > 0 {
		return false
	}
	addSysPost := entity.SysPost{
		PostCode:   sysPost.PostCode,
		PostName:   sysPost.PostName,
		PostStatus: sysPost.PostStatus,
		CreateTime: util.HTime{Time: time.Now()},
		Remark:     sysPost.Remark,
	}
	tx := Db.Save(&addSysPost)
	if tx.RowsAffected > 0 {
		return true
	}
	return false
}
