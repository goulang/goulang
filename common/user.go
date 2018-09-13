package common

import (
	"math/rand"
	"path"
	"time"
)

const (
	Lnormal   = iota // 正常: 正常用户
	Linactive        // 未激活：需要发邮件激活，不能登陆
	Ldisable         // 禁用：不能登陆
	Lnowrite         // 不可写: 不能进行发帖和评论
)

func GetFileUniqueName(name string) string {
	t := time.Now().UnixNano()
	rand.Seed(t)
	r := rand.Int63()
	hash := GetMD5Hash(string(t) + string(name) + string(r))
	base := path.Base(name)
	ext := path.Ext(base)
	return hash+ext
}