package api

type Identity struct {
	theType string
}

/*
func (i *Identity) GetType() string {
	return i.theType
}
*/

type IdentityType interface {
	GetType() string
}

type IdHandler interface {
	GetElements(id IdentityType) map[string]string
	FromElements(map[string]string) IdentityType
}

type Wallet interface {
	Import(label string, id IdentityType) error
	Export(label string) (IdentityType, error)
	Delete(label string) error
	Exists(label string) bool
	//List
}
