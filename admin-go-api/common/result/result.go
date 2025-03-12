//通用返回结构

package result

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 结构体

type Result struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` //提示信息
	Data    interface{} `json:"data"`    //返回的数据

}

// 返回成功

func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Code = int(ApiCode.SUCCESS)
	res.Message = ApiCode.GetMessage(ApiCode.SUCCESS)
	res.Data = data
	c.JSON(http.StatusOK, res)

}

// 返回失败

func Failed(c *gin.Context, code int, message string) {
	res := Result{}
	res.Code = code
	res.Message = message
	c.JSON(http.StatusOK, res)
}
