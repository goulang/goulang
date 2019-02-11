package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/goulang/goulang/models"
	"github.com/goulang/goulang/proxy"
)

// GetTopics get all Topic
func GetTopics(c *gin.Context) {
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

	data, err := proxy.Topic.GetMany(nil, page, limit)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	topics := data.([]models.Topic)

	n, err := proxy.Topic.Count(nil)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(200, map[string]interface{}{
		"page":  page,
		"limit": limit,
		"count": n,
		"list":  topics,
	})
}

// GetTopic get a Topic
func GetTopic(c *gin.Context) {
	topicID := c.Param("topicID")
	data, err := proxy.Topic.Get(topicID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	topic := data.(models.Topic)
	c.JSON(200, topic)

	go func() {
		proxy.Topic.View(topicID)
	}()
}

// CreateTopics create a Topic
func CreateTopic(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	var topic models.Topic
	err := c.BindJSON(&topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	if topic.Title == "" {
		c.String(400, "标题不能为空")
		return
	}

	topic.UserID = user.ID
	err = proxy.Topic.Create(&topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": 0,
		"errmsg":"保存成功",
	})
}

// UpdateTopics update a Topic
func UpdateTopic(c *gin.Context) {
	topicID := c.Param("topicID")
	var topic models.Topic
	err := c.Bind(&topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	err = proxy.Topic.Update(topicID, &topic)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

// DeleteTopics delete a Topic
func DeleteTopic(c *gin.Context) {
	topicID := c.Param("topicID")
	err := proxy.Topic.Delete(topicID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
}
