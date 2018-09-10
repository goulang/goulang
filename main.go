package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/routes"
)

func main() {
	router := gin.Default()

	apiGroup := router.Group("api")

	// auth controller
	authGroup := apiGroup.Group("auth")
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

	router.Run(":" + os.Getenv("PORT"))
}
