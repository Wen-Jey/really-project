//跨域中间件

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-type, AccessToken, X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET,OPTIONS")
		c.Header("Access-Control-Expose-Headers", "true")

		//放行所有OPTION方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)

		}
		//处理访问情况
		c.Next()

	}
}
