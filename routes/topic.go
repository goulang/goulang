package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/models"
)

// GetTopics get all Topic
func GetTopics(c *gin.Context) {
	var Topics []models.Topic
	err := topicCollection.Find(bson.M{}).All(&Topics)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, Topics)
}

// GetTopic get a Topic
func GetTopic(c *gin.Context) {
	TopicID := c.Param("TopicID")
	var Topic models.Topic
	err := topicCollection.FindId(bson.ObjectIdHex(TopicID)).One(&Topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, Topic)
}

// CreateTopics create a Topic
func CreateTopic(c *gin.Context) {
	var Topic models.Topic
	err := c.BindJSON(&Topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	Topic.ID = bson.NewObjectId()
	now := time.Now()
	Topic.CreatedAt = now
	Topic.UpdatedAt = now
	err = topicCollection.Insert(&Topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// UpdateTopics update a Topic
func UpdateTopic(c *gin.Context) {
	TopicID := c.Param("TopicID")
	var Topic models.Topic
	err := c.Bind(&Topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = topicCollection.UpdateId(bson.ObjectIdHex(TopicID), bson.M{
		"$set": Topic,
	})
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// DeleteTopics delete a Topic
func DeleteTopic(c *gin.Context) {
	TopicID := c.Param("TopicID")
	err := topicCollection.RemoveId(bson.ObjectIdHex(TopicID))
	if err != nil {
		c.String(400, err.Error())
		return
	}
}
