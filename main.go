package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/routes"
)

func main() {
	router := gin.New()
	loadMiddlewares(router)
	loadRouters(router)
	router.Run(":" + os.Getenv("PORT"))
}

func loadMiddlewares(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// cors
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	config.AddAllowMethods("PUT", "DELETE")
	r.Use(cors.New(config))

	// session
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Domain: "localhost",
		MaxAge: 3 * 24 * 3600,
	})
	r.Use(sessions.Sessions("goulang", store))
}

func loadRouters(r *gin.Engine) {
	// qiniu
	qiniuGroup := r.Group("qiniu")
	qiniuGroup.GET("token", routes.GetUploadToken)
	// qiniuGroup.POST("callback", routes.CallbackURL)
	qiniuGroup.POST("test", routes.Test)

	// 登录
	r.POST("login", routes.Login)
	// 注销
	r.POST("logout", routes.LoginRequred, routes.Logout)
	// 当前用户
	r.GET("user", routes.LoginRequred, routes.User)
	// 注册(默认已激活)
	r.POST("regist", routes.Regist)
	// 修改密码 TODO
	r.POST("passwd/:userID", routes.LoginRequred, routes.UserOwnerRequred, routes.Passwd)
	// 激活账户 TODO
	r.GET("active", routes.Active)
	// 查看其他用户
	r.GET("users/:userID", routes.GetUser)
	// 修改个人信息 TODO
	r.POST("users/userID", routes.LoginRequred, routes.UserOwnerRequred, routes.UpdateProfile)
	// 上传头像 TODO
	r.POST("avatar/:userID", routes.LoginRequred, routes.UserOwnerRequred, routes.Avatar)
	// 删除用户
	r.DELETE("users/:userID", routes.LoginRequred, routes.UserOwnerRequred, routes.DeleteUser)

	// 帖子列表
	r.GET("topics", routes.GetTopics)
	// 查看帖子
	r.GET("topics/:topicID", routes.GetTopic)
	// 发帖
	r.POST("topics", routes.LoginRequred, routes.CreateTopic)
	// 修改帖子
	r.PUT("topics/:topicID", routes.LoginRequred, routes.TopicOwnerRequred, routes.UpdateTopic)
	// 删帖
	r.DELETE("topics/:topicID", routes.LoginRequred, routes.TopicOwnerRequred, routes.DeleteTopic)

	// 获取评论
	r.GET("topics/:topicID/comments", routes.GetComments)
	// 发表评论
	r.POST("topics/:topicID/comments", routes.LoginRequred, routes.CreateComment)
	// 删除评论
	r.DELETE("topics/:topicID/comments/:commentID", routes.LoginRequred, routes.CommentOwnerRequred, routes.DeleteComment)
}
