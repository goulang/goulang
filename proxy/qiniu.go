package proxy

var _ proxy = &qiniuProxy{}

type qiniuProxy struct {
	baseProxy
}
