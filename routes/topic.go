package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/wodog/goulang/models"
// 	"github.com/wodog/goulang/proxy"
// )

// // GetTopics get all Topic
// func GetTopics(c *gin.Context) {
// 	data, err := proxy.Topic.GetMany(nil, 1, 10)
// 	if err != nil {
// 		c.String(400, err.Error())
// 		return
// 	}
// 	c.JSON(200, data)
// }

// // GetTopic get a Topic
// func GetTopic(c *gin.Context) {
// 	TopicID := c.Param("TopicID")
// 	data, err := proxy.Topic.Get(TopicID)
// 	if err != nil {
// 		c.String(400, err.Error())
// 		return
// 	}
// 	c.JSON(200, data)
// }

// // CreateTopics create a Topic
// func CreateTopic(c *gin.Context) {
// 	var topic models.Topic
// 	err := c.BindJSON(&topic)
// 	if err != nil {
// 		c.String(400, err.Error())
// 		return
// 	}
// 	err = proxy.Topic.Create(&topic)
// 	if err != nil {
// 		c.String(400, err.Error())
// 		return
// 	}
// }

// // // UpdateTopics update a Topic
// // func UpdateTopic(c *gin.Context) {
// // 	TopicID := c.Param("TopicID")
// // 	var Topic models.Topic
// // 	err := c.Bind(&Topic)
// // 	if err != nil {
// // 		c.String(400, err.Error())
// // 		return
// // 	}
// // 	err = topicCollection.UpdateId(bson.ObjectIdHex(TopicID), bson.M{
// // 		"$set": Topic,
// // 	})
// // 	if err != nil {
// // 		c.String(400, err.Error())
// // 		return
// // 	}
// // }

// // // DeleteTopics delete a Topic
// // func DeleteTopic(c *gin.Context) {
// // 	TopicID := c.Param("TopicID")
// // 	err := topicCollection.RemoveId(bson.ObjectIdHex(TopicID))
// // 	if err != nil {
// // 		c.String(400, err.Error())
// // 		return
// // 	}
// // }
