package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/models"
)

// GetUsers get all user
func GetUsers(c *gin.Context) {
	var users []models.User
	err := userCollection.Find(bson.M{}).All(&users)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, users)
}

// GetUser get a user
func GetUser(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User
	err := userCollection.FindId(bson.ObjectIdHex(userID)).One(&user)
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
	user.ID = bson.NewObjectId()
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
	err = userCollection.Insert(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// UpdateUsers update a user
func UpdateUser(c *gin.Context) {
	userID := c.Param("userID")
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = userCollection.UpdateId(bson.ObjectIdHex(userID), bson.M{
		"$set": user,
	})
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// DeleteUsers delete a user
func DeleteUser(c *gin.Context) {
	userID := c.Param("userID")
	err := userCollection.RemoveId(bson.ObjectIdHex(userID))
	if err != nil {
		c.String(400, err.Error())
		return
	}
}
