package api

import "errors"

/*
type Wallet interface {
	Import(Identity) error
    Export(label string) (Identity, error)
    Delete(label string) error
    Exists(label string) error
    //List
}
*/

type InMemoryWallet struct {
	idhandler IdHandler
	storage   map[string]map[string]string
}

func NewInMemoryWallet(idhandler IdHandler) *InMemoryWallet {
	return &InMemoryWallet{idhandler, make(map[string]map[string]string, 10)}
}

func (f *InMemoryWallet) Import(label string, id IdentityType) error {
	elements := f.idhandler.GetElements(id)
	f.storage[label] = elements
	return nil
}

func (f *InMemoryWallet) Export(label string) (IdentityType, error) {
	if elements, ok := f.storage[label]; ok {
		return f.idhandler.FromElements(elements), nil
	}
	return nil, errors.New("label doesn't exist")
}

func (f *InMemoryWallet) Delete(label string) error {
	if _, ok := f.storage[label]; ok {
		delete(f.storage, label)
		return nil
	}
	return nil // what should we do here ?
}

func (f *InMemoryWallet) Exists(label string) bool {
	_, ok := f.storage[label]
	return ok
}
