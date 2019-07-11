package api

type X509Identity struct {
	Identity
	cert string
	key  string
}

func (x *X509Identity) GetType() string {
	return "Identity.theType"
}

func (x *X509Identity) GetCert() string {
	return x.cert
}

func (x *X509Identity) GetKey() string {
	return x.key
}

func NewX509Identity(cert string, key string) *X509Identity {
	return &X509Identity{Identity{"X509"}, cert, key}
}

type X509IdentityHandler struct{}

func (x *X509IdentityHandler) GetElements(id IdentityType) map[string]string {
	r, _ := id.(*X509Identity)

	return map[string]string{
		"cert": r.cert,
		"key":  r.key,
	}
}

func (x *X509IdentityHandler) FromElements(elements map[string]string) IdentityType {
	y := &X509Identity{Identity{"X509"}, elements["cert"], elements["key"]}
	return y
}

func NewX509IdentityHandler() *X509IdentityHandler {
	return &X509IdentityHandler{}
}
