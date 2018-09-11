package proxy

var _ proxy = &topicProxy{}

type topicProxy struct {
	baseProxy
}
