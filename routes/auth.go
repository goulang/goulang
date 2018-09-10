package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/models"
)

func Info(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	user = user.(models.User)
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = userCollection.Find(bson.M{
		"name":     user.Name,
		"password": user.Password,
	}).One(&user)
	if err != nil {
		c.String(400, "用户名或者密码不正确")
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

// THINK: 与创建用户接口重复了, 考虑是否保留
func Regist(c *gin.Context) {

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
