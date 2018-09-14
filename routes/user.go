package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/common"
	"github.com/goulang/goulang/errors"
	"github.com/goulang/goulang/models"
	"github.com/goulang/goulang/proxy"
	"github.com/goulang/goulang/storage/Qiniu"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	data, err := proxy.User.GetOne(bson.M{
		"name":     user.Name,
		"password": common.GetMD5Hash(user.Password),
	})
	if err != nil {
		ApiStandardError := errors.ApiErrNamePwdIncorrect
		c.JSON(400, ApiStandardError)
		return
	}
	user = data.(models.User)

	// 校验账号状态
	if user.Status == common.Linactive || user.Status == common.Ldisable {
		c.String(403, "账户禁止登陆")
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
	session.Clear()
	err := session.Save()
	if err != nil {
		c.String(400, err.Error())
		return
	}
}

func User(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	c.JSON(200, user)
}

func Regist(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	user.Status = common.Lnormal
	user.Password = common.GetMD5Hash(user.Password)
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
	data, err := proxy.User.GetMany(nil, 1, 10)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	users := data.([]models.User)
	c.JSON(200, users)
}

// GetUser get a user
func GetUser(c *gin.Context) {
	userID := c.Param("userID")
	data, err := proxy.User.Get(userID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	user := data.(models.User)
	c.JSON(200, user)
}

func Passwd(c *gin.Context) {
	var user models.User
	var password models.Password
	userID := c.Param("userID")
	err := c.BindJSON(&password)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}
	data, err := proxy.User.Get(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}
	user = data.(models.User)

	if !common.CheckPasswrd(password.Password, user.Password) {
		c.JSON(http.StatusBadRequest, errors.ApiErrPwdIncorrect)
		return
	}

	user.Password = common.GetMD5Hash(password.RePassword)
	if err := proxy.User.Update(userID, &user); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}

	c.JSON(http.StatusOK, errors.ApiErrSuccess)
}

func Active(c *gin.Context) {

}

func UpdateProfile(c *gin.Context) {
	userID := c.Param("userID")
	var update models.Update
	err := c.BindJSON(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}

	if err := proxy.User.Update(userID, &update); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}

	c.JSON(http.StatusOK, errors.ApiErrSuccess)
}

func Avatar(c *gin.Context) {
	var user models.User
	userID := c.Param("userID")
	data, err := proxy.User.Get(userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}
	user = data.(models.User)

	//删除原有头像
	ok := Qiniu.Storage.DeleteFile(user.Avatar, true)
	if ok != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}

	//上传新头像
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}
	name := common.GetFileUniqueName(header.Filename)
	// TODO 完成从配置读取路径
	name = time.Now().Format("avatar/2006/01/02") + "/" + name
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}
	_, isOk := Qiniu.Storage.PutFile(name, bytes)
	if isOk != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}

	//更新头像
	user.Avatar = Qiniu.Storage.GetUrl(name)
	if err := proxy.User.Update(userID, &user); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, errors.NewUnknownErr(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"key":    name,
		"access": Qiniu.Storage.GetUrl(name),
	})
}
