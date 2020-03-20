package routes

import (
	"gin-admin/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.StaticFS("/statics1", http.Dir("website/AdminLTE-2.4.12/bower_components"))
	router.StaticFS("/statics2", http.Dir("website/AdminLTE-2.4.12/dist"))
	router.StaticFS("/statics3", http.Dir("website/AdminLTE-2.4.12/plugins"))
	router.LoadHTMLGlob("website/AdminLTE-2.4.12/pages/**/*.html")

	user := router.Group("")
	{
		User := new(controllers.User)
		user.GET("/login", User.Login)
	}

	web := router.Group("")
	{
		adminUser := new(controllers.AdminUser)
		web.GET("/test", adminUser.Test)
		web.GET("/session_test", adminUser.SessionTest)
		web.GET("/template_test", adminUser.Template)
	}

	v1 := router.Group("/api/v1")
	{
		adminUser := new(controllers.AdminUser)
		v1.GET("/admin-users", adminUser.Index)
		v1.POST("/admin-users", adminUser.Store)
		v1.PATCH("/admin-users/:id", adminUser.Update)
		v1.DELETE("/admin-users/:id", adminUser.Destroy)
		v1.GET("/admin-users/:id", adminUser.Show)
	}

	return router

}
