package routes

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/models"
	"github.com/goulang/goulang/proxy"
)

// GetComments get all Comment
func GetComments(c *gin.Context) {
	topicID := c.Param("topicID")
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.String(400, err.Error())
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.String(400, err.Error())
		return
	}

	data, err := proxy.Comment.GetMany(bson.M{
		"topic_id": bson.ObjectIdHex(topicID),
	}, page, limit)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	comments := data.([]models.Comment)

	n, err := proxy.Comment.Count(nil)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(200, map[string]interface{}{
		"page":  page,
		"limit": limit,
		"count": n,
		"list":  comments,
	})
}

// CreateComments create a Comment
func CreateComment(c *gin.Context) {
	topicID := c.Param("topicID")
	user := c.MustGet("user").(models.User)

	var comment models.Comment
	err := c.BindJSON(&comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	if comment.Content == "" {
		c.String(400, "内容不能为空")
		return
	}

	comment.UserID = user.ID
	comment.TopicID = bson.ObjectIdHex(topicID)
	err = proxy.Comment.Create(&comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// DeleteComments delete a Comment
func DeleteComment(c *gin.Context) {
	CommentID := c.Param("CommentID")
	err := proxy.Comment.Delete(CommentID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}
