package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	// "github.com/gin-contrib/cors"
	"net/http"

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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "http://localhost:9000")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-requested-with")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			// OPTIONS请求直接返回204
			c.AbortWithStatus(http.StatusNoContent)
 
		}

		// 处理请求
		c.Next()
	}
}
func loadMiddlewares(r *gin.Engine) {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(Cors())
	// cors
	// config := cors.DefaultConfig()
	// config.AllowCredentials = true
	// config.AllowOriginFunc = func(origin string) bool {
	// 	return true
	// }
	// config.AddAllowMethods("POST", "PUT", "DELETE")
	// r.Use(cors.New(config))

	// session
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		Domain: "goulang.org",
		MaxAge: 3 * 24 * 3600,
	})
	r.Use(sessions.Sessions("goulang", store))
}

func loadRouters(r *gin.Engine) {
	// qiniu
	qiniuGroup := r.Group("qiniu")
	//获取上传令牌
	qiniuGroup.GET("token", routes.GetUploadToken)
	//前端上传回调地址
	qiniuGroup.POST("callback", routes.CallbackURL)
	//测试使用
	//qiniuGroup.POST("test", routes.Test)

	// 登录
	r.POST("login", routes.Login)
	// 注销
	r.POST("logout", routes.LoginRequred, routes.Logout)
	// 当前用户
	r.GET("user", routes.LoginRequred, routes.User)
	// 注册(默认已激活)
	r.POST("register", routes.Regist)
	// 修改密码
	r.POST("passwd/:userID", routes.LoginRequred, routes.UserOwnerRequred, routes.Passwd)
	// 激活账户
	r.GET("active/:active", routes.Active)
	// 查看其他用户
	r.GET("users/:userID", routes.GetUser)
	// 修改个人信息
	r.POST("users/:userID", routes.UpdateProfile)
	// 上传头像
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
	r.GET("/", routes.All)
}
