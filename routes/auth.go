package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/models"
	"github.com/goulang/goulang/proxy"
)

// LoginRequred 需要登陆
func LoginRequred(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.AbortWithStatus(401)
		return
	}
	c.Set("user", user)
}

// UserRequred 需要当前用户或者管理员
func UserOwnerRequred(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)
	if currentUser.Admin == true {
		return
	}

	userID := c.Param("userID")
	Iuser, err := proxy.User.Get(userID)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	user := Iuser.(models.User)
	if user.ID != currentUser.ID {
		c.AbortWithStatus(403)
		return
	}
}

// TopicRequred 需要当前帖子用户或者管理员
func TopicOwnerRequred(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)
	if currentUser.Admin == true {
		return
	}

	topicID := c.Param("topicID")
	Itopic, err := proxy.Topic.Get(topicID)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	topic := Itopic.(models.Topic)
	if topic.UserID != currentUser.ID {
		c.AbortWithStatus(403)
		return
	}
}

// CommentRequred 需要当前帖子用户或者管理员
func CommentOwnerRequred(c *gin.Context) {
	currentUser := c.MustGet("user").(models.User)
	if currentUser.Admin == true {
		return
	}

	CommentID := c.Param("CommentID")
	IComment, err := proxy.Comment.Get(CommentID)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	comment := IComment.(models.Comment)
	if comment.UserID != currentUser.ID {
		c.AbortWithStatus(403)
		return
	}
}

// AdminRequred 需要管理员
func AdminRequred(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	if user.Admin == true {
	} else {
		c.AbortWithStatus(403)
		return
	}
}
