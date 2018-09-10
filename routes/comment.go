package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/models"
)

// GetComments get all Comment
func GetComments(c *gin.Context) {
	var Comments []models.Comment
	err := commentCollection.Find(bson.M{}).All(&Comments)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, Comments)
}

// GetComment get a Comment
func GetComment(c *gin.Context) {
	CommentID := c.Param("CommentID")
	var Comment models.Comment
	err := commentCollection.FindId(bson.ObjectIdHex(CommentID)).One(&Comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, Comment)
}

// CreateComments create a Comment
func CreateComment(c *gin.Context) {
	var Comment models.Comment
	err := c.BindJSON(&Comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	Comment.ID = bson.NewObjectId()
	now := time.Now()
	Comment.CreatedAt = now
	Comment.UpdatedAt = now
	err = commentCollection.Insert(&Comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// UpdateComments update a Comment
func UpdateComment(c *gin.Context) {
	CommentID := c.Param("CommentID")
	var Comment models.Comment
	err := c.Bind(&Comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = commentCollection.UpdateId(bson.ObjectIdHex(CommentID), bson.M{
		"$set": Comment,
	})
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// DeleteComments delete a Comment
func DeleteComment(c *gin.Context) {
	CommentID := c.Param("CommentID")
	err := commentCollection.RemoveId(bson.ObjectIdHex(CommentID))
	if err != nil {
		c.String(400, err.Error())
		return
	}
}
