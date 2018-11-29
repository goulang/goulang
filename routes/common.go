package routes

import "github.com/gin-gonic/gin"

func All (c *gin.Context){
	c.JSON(200, "OK")
}