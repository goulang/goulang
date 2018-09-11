package routes

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/goulang/goulang/errors"
	"github.com/goulang/goulang/models"
	"github.com/globalsign/mgo/bson"
	"time"
	"github.com/qiniu/api.v7/storage"
	"fmt"
	"github.com/qiniu/x/rpc.v7"
	"log"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	bucket    = os.Getenv("QINIU_TEST_BUCKET")
)

/*
获取上传Token
*/
func GetUploadToken(c *gin.Context) {

	putPolicy := storage.PutPolicy{
		Scope:            bucket,
		Expires:          7200,
		CallbackURL:      "http://esp.juneandlee.cn/api/redirect",
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"}`,
		CallbackBodyType: "application/json",
	}
	upToken := putPolicy.UploadToken(getMac())
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
	//完成七牛回调验证
	isQiniu, err := qbox.VerifyCallback(getMac(), c.Request)
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
	keys := []string{
		"1.jpg",
		"2.jpg",
		"3.jpg",
		"4.jpg",
	}

	BatchDeleteFile(keys)
}

//获取文件信息 传入文件KEY
func FileInfo(key string) storage.FileInfo {
	manager := newBucketManager()
	fileInfo, sErr := manager.Stat(bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return storage.FileInfo{}
	}
	return fileInfo
}

//获取指定前缀列表文件 (前缀)prefix (分隔符)delimiter (标记)marker (长度)limit
func PrefixListFiles(limit int) (data [][]storage.ListItem) {
	delimiter := ""
	prefix := ""
	marker := ""
	manager := newBucketManager()
	for {
		entries, _, nextMarker, hashNext, err := manager.ListFiles(bucket, prefix, delimiter, marker, limit)
		if err != nil {
			fmt.Println("list error,", err)
			break
		}
		//print entries
		data = append(data, entries)
		if hashNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
	return
}

//修改文件MimeType 传入 (文件名)key (新的Mine) newMine
func ChangeMimeType(key string, newMime string) error {
	manager := newBucketManager()
	err := manager.ChangeMime(bucket, key, newMime)
	if err != nil {
		return err
	}
	return nil
}

//删除空间中的文件 (文件名)key
func DeleteFile(key string) error {
	manager := newBucketManager()
	err := manager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//批量删除
func BatchDeleteFile(keys []string) error {
	manager := newBucketManager()

	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	}

	rets, err := manager.Batch(deleteOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*rpc.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				log.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					log.Printf("%s\n", ret.Data.Error)
					return err
				}
			}
		} else {
			return err
		}
	}
	return nil
}

func getMac() *qbox.Mac {
	return qbox.NewMac(accessKey, secretKey)
}

func newBucketManager() *storage.BucketManager {
	cfg := storage.Config{
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(getMac(), &cfg)
	return bucketManager
}
