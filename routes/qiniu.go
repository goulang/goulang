package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/storage/Qiniu"
	"io/ioutil"
)


/*
获取上传Token
*/
func GetUploadToken(c *gin.Context) {
	upToken, putPolicy := Qiniu.Storage.GetUploadToken()
	c.JSON(200, gin.H{
		"token":   upToken,
		"expires": putPolicy.Expires,
	})
	return
}

/*
回调保存上传信息
*/
func CallbackURL(c *gin.Context) {
	// //完成七牛回调验证
	// isQiniu, err := qbox.VerifyCallback(storage.Mac, c.Request)
	// if err != nil {
	// 	c.JSON(200, errors.NewUnknownErr(err))
	// 	return
	// }
	// if !isQiniu {
	// 	c.JSON(200, errors.ApiErrRefuse)
	// 	return
	// }

	// //入库
	// var file models.QFile
	// if err := c.BindJSON(&file); err != nil {
	// 	c.JSON(200, errors.NewUnknownErr(err))
	// 	return
	// }
	// file.ID = bson.NewObjectId()
	// now := time.Now()
	// file.CreatedAt = now
	// file.UpdatedAt = now

	// if err := qiniuCollection.Insert(&file); err != nil {
	// 	c.JSON(200, errors.NewUnknownErr(err))
	// 	return
	// }
	// c.JSON(200, file)
	// return
}

func Test(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println("file", err)
	}
	bytes, _ := ioutil.ReadAll(file)
	rest, _ := Qiniu.Storage.PutFile(header.Filename, bytes)
	fmt.Println(rest)
	return
}

