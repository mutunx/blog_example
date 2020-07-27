package routers

import (
	v1 "ginExample/routers/api/v1"
	"github.com/gin-gonic/gin"

	setting "ginExample/pkg"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/tags", v1.GetTags)
		v1Group.POST("/tags", v1.AddTags)
		v1Group.PUT("/tags/:1", v1.EditTags)
		v1Group.DELETE("/tags/:id", v1.DeleteTags)
	}

	return r
}
