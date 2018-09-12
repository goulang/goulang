package common

const (
	Lnormal   = iota // 正常: 正常用户
	Linactive        // 未激活：需要发邮件激活，不能登陆
	Ldisable         // 禁用：不能登陆
	Lnowrite         // 不可写: 不能进行发帖和评论
)
