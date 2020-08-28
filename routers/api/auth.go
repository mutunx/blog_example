package api

import (
	"ginExample/models"
	"ginExample/pkg/e"
	"ginExample/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	// 验证密码
	isExist := models.CheckAuth(username, password)
	if isExist {
		// 获取token
		token, er := util.GenerateToken(username, password)
		if er != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token

			code = e.SUCCESS
		}

	} else {
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}
