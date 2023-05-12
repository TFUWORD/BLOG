package routers

import (
	"web/api"
	"web/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("articles", app.ArticleListView)
	router.GET("articles/:id", app.ArticleDetailView)
}
