package proxy

var _ proxy = &commentProxy{}

type commentProxy struct {
	baseProxy
}
