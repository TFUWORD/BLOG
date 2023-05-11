package routers

import "web/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)
	router.GET("images", app.ImageListView)
	router.GET("image_names", app.ImageNameListView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)
}
