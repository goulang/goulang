package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/wodog/goulang/routes"
)

func loadMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// cors
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	config.AddAllowMethods("PUT", "DELETE")
	router.Use(cors.New(config))

	// session
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Domain: "goulang.com",
		MaxAge: 3 * 24 * 3600,
	})
	router.Use(sessions.Sessions("goulang", store))
}

func main() {
	router := gin.New()
	loadMiddlewares(router)

	// 登录
	router.POST("login", routes.Login)
	// 注销
	router.POST("logout", routes.Logout)
	// 当前用户
	router.GET("user", routes.User)
	// 注册
	router.POST("regist", routes.Regist)
	// 修改密码
	router.POST("passwd", routes.Passwd)
	// 激活账户
	router.GET("active", routes.Active)
	// 查看其他用户
	router.GET("users/:userID", routes.Profile)
	// 修改个人信息
	router.POST("users/:userID", router.UpdateProfile)

	// 帖子列表
	router.GET("topics", router.GetTopics)
	// 查看帖子
	router.GET("topics/:topicID", router.GetTopic)
	// 发帖
	router.POST("topics", router.CreateTopic)
	// 修改帖子
	router.PUT("topics/:topicID", router.UpdateTopic)
	// 删帖
	router.POST("post/:topicID", router.DeleteTopic)

	// // // user controller
	// userGroup := apiGroup.Group("users")
	// userGroup.GET(":userID", routes.GetUser)
	// userGroup.POST("", routes.CreateUser)
	// userGroup.PUT(":userID", routes.UpdateUser)
	// userGroup.DELETE(":userID", routes.DeleteUser)

	// // topic controller
	// topicGroup := apiGroup.Group("topics")
	// topicGroup.GET("", routes.GetTopics)
	// topicGroup.GET(":topicID", routes.GetTopic)
	// topicGroup.POST("", routes.CreateTopic)
	// topicGroup.PUT(":topicID", routes.UpdateTopic)
	// topicGroup.DELETE(":topicID", routes.DeleteTopic)

	// // comment controller
	// commentGroup := apiGroup.Group("comments")
	// commentGroup.GET("", routes.GetComments)
	// commentGroup.GET(":commentID", routes.GetComment)
	// commentGroup.POST("", routes.CreateComment)
	// commentGroup.PUT(":commentID", routes.UpdateComment)
	// commentGroup.DELETE(":commentID", routes.DeleteComment)

	router.Run(":" + os.Getenv("PORT"))
}
