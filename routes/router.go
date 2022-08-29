package routes

import (
	v1 "blog/api/v1"
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//User模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)

		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		//上传文件
		auth.POST("upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("categorys", v1.GetCates)
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)

	}
	r.Run(utils.HttpPort)
}
