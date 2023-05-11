package routers

import (
	"web/api"
	"web/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))


func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("register", app.UserCreateView)
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAdmin(), app.UserUpdateRoleView)
	// router.POST("login", app.QQLoginView())
	
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
	router.POST("logout", middleware.JwtAuth(), app.LogoutView)
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)
	router.POST("user_bing_email", middleware.JwtAuth(), app.UserBindEmailView)
}
