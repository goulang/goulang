package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/goulang/goulang/errors"
	"github.com/goulang/goulang/models"
	"github.com/goulang/goulang/storage/Qiniu"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"os"
	"time"
)

var (
	bucket      string
	accessKey   string
	secretKey   string
	callBackURL string
)

func init() {
	bucket = os.Getenv("QINIU_TEST_BUCKET")
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	callBackURL = os.Getenv("QINIU_CALLBACK_URL")
}

/*
获取上传Token
*/
func GetUploadToken(c *gin.Context) {
	storage := Qiniu.NewQiniu(bucket, accessKey, secretKey, callBackURL)
	c.JSON(200, storage.GetUploadToken())
	return
}

/*
回调保存上传信息
*/
func CallbackURL(c *gin.Context) {
	storage := Qiniu.NewQiniu(bucket, accessKey, secretKey, callBackURL)
	//完成七牛回调验证
	isQiniu, err := qbox.VerifyCallback(storage.Mac, c.Request)
	if err != nil {
		c.JSON(200, errors.NewUnknownErr(err))
		return
	}
	if !isQiniu {
		c.JSON(200, errors.ApiErrRefuse)
		return
	}

	//入库
	var file models.QFile
	if err := c.BindJSON(&file); err != nil {
		c.JSON(200, errors.NewUnknownErr(err))
		return
	}
	file.ID = bson.NewObjectId()
	now := time.Now()
	file.CreatedAt = now
	file.UpdatedAt = now

	if err := qiniuCollection.Insert(&file); err != nil {
		c.JSON(200, errors.NewUnknownErr(err))
		return
	}
	c.JSON(200, file)
	return
}

func Test(c *gin.Context) {
	//storage := Qiniu.NewQiniu(bucket, accessKey, secretKey, callBackURL)
	fmt.Println(accessKey, secretKey)

	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	fileInfo, sErr := bucketManager.Stat(bucket, "FjMFlKUIbFSm1sJhDNY2IJ7OqT3x")
	if sErr != nil {
		fmt.Println(sErr)

	}
	fmt.Println(fileInfo)
	return
}
