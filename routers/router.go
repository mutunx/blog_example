package routers

import (
	_ "ginExample/docs"
	"ginExample/middlewave/jwt"
	setting "ginExample/pkg"
	"ginExample/routers/api"
	v1 "ginExample/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)

	v1Group := r.Group("/api/v1")
	v1Group.Use(jwt.JWT())
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
