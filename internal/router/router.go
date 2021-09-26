package router

import (
	"github.com/gin-gonic/gin"
	"practice.com/blog_service/internal/router/api"
	v1 "practice.com/blog_service/internal/router/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	pong := api.NewPong()
	r.GET("/ping", pong.Pong)

	tagV1 := v1.NewTag()
	articleV1 := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tagV1.Create)
		apiv1.DELETE("/tags/:id", tagV1.Delete)
		apiv1.PUT("/tags/:id", tagV1.Update)
		apiv1.PATCH("/tags/:id/state", tagV1.Update)
		apiv1.GET("/tags", tagV1.List)

		apiv1.POST("/articles", articleV1.Create)
		apiv1.DELETE("/articles/:id", articleV1.Delete)
		apiv1.PUT("/articles/:id", articleV1.Update)
		apiv1.PATCH("/articles/:id/state", articleV1.Update)
		apiv1.GET("/articles/:id", articleV1.Get)
		apiv1.GET("/articles", articleV1.List)
	}

	return r
}
