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
		v1Group.PUT("/tags/:id", v1.EditTags)
		v1Group.DELETE("/tags/:id", v1.DeleteTags)
		v1Group.GET("/articles", v1.GetArticles)
		v1Group.GET("/articles/:id", v1.GetArticle)
		v1Group.POST("/articles", v1.AddArticle)
		v1Group.PUT("/articles/:id", v1.EditArticle)
		v1Group.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
