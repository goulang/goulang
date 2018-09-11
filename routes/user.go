package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/models"
	"github.com/goulang/goulang/proxy"
)

// GetUsers get all user
func GetUsers(c *gin.Context) {
	users, err := proxy.User.GetMany(nil, 1, 10)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, users)
}

// GetUser get a user
func GetUser(c *gin.Context) {
	userID := c.Param("userID")
	user, err := proxy.User.Get(userID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, user)
}

// CreateUsers create a user
func CreateUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	err = proxy.User.Create(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// UpdateUsers update a user
func UpdateUser(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = proxy.User.Update(userID, &user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// DeleteUsers delete a user
func DeleteUser(c *gin.Context) {
	userID := c.Param("userID")
	err := proxy.User.Delete(userID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}
