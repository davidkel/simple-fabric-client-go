package api

import "fmt"

type NetworkOptions struct {
	Wallet   Wallet
	Identity string
}

type NetworkInterface interface {
	Init(ccp string, networkOptions *NetworkOptions) error
	Dispose() error
	//GetContext()
	GetLedger(channelName string) (Ledger, error)
}

type Network struct {
	id IdentityType
	// cache of the ledgers
}

func (network *Network) Init(ccp string, networkOptions *NetworkOptions) error {
	id, _ := networkOptions.Wallet.Export(networkOptions.Identity)
	fmt.Println(id)
	return nil
}

func (network *Network) Dispose() error {
	return nil
}

func (network *Network) GetLedger(channelName string) (*Ledger, error) {
	return newLedger(), nil
}

func NewNetwork() Network {
	return Network{}
}
