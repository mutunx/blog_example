package jwt

import (
	"ginExample/pkg/e"
	"ginExample/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		// 设置默认值
		code = e.SUCCESS
		// token处理
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claim, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			// 阻止其他中间件调用
			c.Abort()
			return
		}
		c.Next()
	}
}
