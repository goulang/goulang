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
	apiGroup := router.Group("api")

	// auth controller
	authGroup := apiGroup.Group("auth")
	authGroup.GET("info", routes.Info)
	authGroup.POST("login", routes.Login)
	authGroup.POST("regist", routes.Regist)
	authGroup.POST("logout", routes.Logout)

	// user controller
	userGroup := apiGroup.Group("users")
	userGroup.GET("", routes.GetUsers)
	userGroup.GET(":userID", routes.GetUser)
	userGroup.POST("", routes.CreateUser)
	userGroup.PUT(":userID", routes.UpdateUser)
	userGroup.DELETE(":userID", routes.DeleteUser)

	// topic controller
	topicGroup := apiGroup.Group("topics")
	topicGroup.GET("", routes.GetTopics)
	topicGroup.GET(":topicID", routes.GetTopic)
	topicGroup.POST("", routes.CreateTopic)
	topicGroup.PUT("/:topicID", routes.UpdateTopic)
	topicGroup.DELETE("/:topicID", routes.DeleteTopic)

	// comment controller
	commentGroup := apiGroup.Group("comments")
	commentGroup.GET("", routes.GetComments)
	commentGroup.GET(":commentID", routes.GetComment)
	commentGroup.POST("", routes.CreateComment)
	commentGroup.PUT(":commentID", routes.UpdateComment)
	commentGroup.DELETE(":commentID", routes.DeleteComment)

	// qiniu controller
	qiniuGroup := apiGroup.Group("qiniu")
	qiniuGroup.GET("token", routes.GetUploadToken)
	qiniuGroup.POST("callback", routes.CallbackURL)
	qiniuGroup.POST("test", routes.Test)

	router.Run(":" + os.Getenv("PORT"))
}
