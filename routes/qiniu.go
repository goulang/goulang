package routes

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/storage/Qiniu"
)

var (
	storage     *Qiniu.Qiniu
	bucket      = os.Getenv("QINIU_TEST_BUCKET")
	accessKey   = os.Getenv("QINIU_ACCESS_KEY")
	secretKey   = os.Getenv("QINIU_SECRET_KEY")
	callBackURL = os.Getenv("QINIU_CALLBACK_URL")
)

func init() {
	storage = Qiniu.NewQiniu(bucket, accessKey, secretKey, callBackURL)
}

/*
获取上传Token
*/
func GetUploadToken(c *gin.Context) {
	upToken, putPolicy := storage.GetUploadToken()
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
	bytes, err := ioutil.ReadAll(file)
	storage.PutFile(header.Filename, bytes)
	return
}
