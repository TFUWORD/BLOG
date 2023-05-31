package routers

import (
	"web/api"
	"web/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.POST("articles/collect", middleware.JwtAuth(), app.ArticleCollCreateView)

	router.PUT("articles", app.ArticleUpdateView)
	router.DELETE("articles", app.ArticleRemoveView)
	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	
	router.GET("articles", app.ArticleListView)
	router.GET("articles/collect", middleware.JwtAuth(), app.ArticleCollListView)
	router.GET("articles/tags", app.ArticleTagListView)
	router.GET("articles/calendar", app.ArticleCalendarView)
	router.GET("articles/:id", app.ArticleDetailView)
	router.GET("articles/detial", app.ArticleDetailByTitleView)
	
}
