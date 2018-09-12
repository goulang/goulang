package proxy

var _ proxy = &userProxy{}

type userProxy struct {
	baseProxy
}
