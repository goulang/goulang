package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/wodog/goulang/errors"
	"github.com/wodog/goulang/models"
	"github.com/wodog/goulang/proxy"
)

func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	success := proxy.User.Login(user.Name, user.Password)
	if !success {
		ApiStandardError := errors.ApiErrNamePwdIncorrect
		c.JSON(400, ApiStandardError)
		return
	}
	session := sessions.Default(c)
	session.Set("user", user)
	err = session.Save()
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	err := session.Save()
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

func User(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	user = user.(models.User)
	c.JSON(200, user)
}

func Regist(c *gin.Context) {
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

// DeleteUsers delete a user
func DeleteUser(c *gin.Context) {
	userID := c.Param("userID")
	err := proxy.User.Delete(userID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// GetUsers get all user
func Users(c *gin.Context) {
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

func Passwd(c *gin.Context) {

}

func Active(c *gin.Context) {

}

func Profile(c *gin.Context) {

}

func UpdateProfile(c *gin.Context) {

}
