package Qiniu

import (
	"fmt"
	"github.com/gin-gonic/gin"
	finalFileInfo "github.com/goulang/goulang/storage"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/x/rpc.v7"
	"log"
)

type Qiniu struct {
	Mac           *qbox.Mac
	bucket        string
	accessKey     string
	secretKey     string
	callBackURL   string
	BucketManager *storage.BucketManager
}

func NewQiniu(bucket, accessKey, secretKey, callBackURL string) *Qiniu {
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: false,
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	return &Qiniu{
		Mac:           mac,
		bucket:        bucket,
		accessKey:     accessKey,
		secretKey:     secretKey,
		callBackURL:   callBackURL,
		BucketManager: bucketManager,
	}
}

/*
获取上传Token
*/
func (q *Qiniu) GetUploadToken() map[string]interface{} {
	putPolicy := storage.PutPolicy{
		Scope:            q.bucket,
		Expires:          7200,
		CallbackURL:      q.callBackURL,
		CallbackBody:     `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)"}`,
		CallbackBodyType: "application/json",
	}
	upToken := putPolicy.UploadToken(q.Mac)

	return gin.H{
		"token":   upToken,
		"expires": putPolicy.Expires,
	}
}

//获取文件信息 传入文件KEY
func (q *Qiniu) FileInfo(key string) finalFileInfo.FinalFileInfo {
	fileInfo, sErr := q.BucketManager.Stat(q.bucket, key)
	if sErr != nil {
		fmt.Println(sErr)
		return finalFileInfo.FinalFileInfo{}
	}
	return finalFileInfo.FinalFileInfo{fileInfo}
}

//获取指定前缀列表文件 (前缀)prefix (分隔符)delimiter (标记)marker (长度)limit
func (q *Qiniu) PrefixListFiles(prefix string, limit int) (data [][]storage.ListItem) {
	delimiter := ""
	marker := ""
	manager := q.newBucketManager()
	for {
		entries, _, nextMarker, hashNext, err := manager.ListFiles(q.bucket, prefix, delimiter, marker, limit)
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
func (q *Qiniu) ChangeMimeType(key string, newMime string) error {
	manager := q.newBucketManager()
	err := manager.ChangeMime(q.bucket, key, newMime)
	if err != nil {
		return err
	}
	return nil
}

//删除空间中的文件 (文件名)key
func (q *Qiniu) DeleteFile(key string) error {
	manager := q.newBucketManager()
	err := manager.Delete(q.bucket, key)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//批量删除
func (q *Qiniu) BatchDeleteFile(keys []string) error {
	manager := q.newBucketManager()

	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(q.bucket, key))
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

func (q *Qiniu) newBucketManager() *storage.BucketManager {
	cfg := storage.Config{
		UseHTTPS: true,
	}
	bucketManager := storage.NewBucketManager(q.Mac, &cfg)
	return bucketManager
}
