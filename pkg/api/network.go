package api

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	othermsp "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

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
	id  IdentityType
	sdk *fabsdk.FabricSDK
	si  othermsp.SigningIdentity
	// cache of the ledgers
}

func (network *Network) Init(ccp string, networkOptions *NetworkOptions) error {

	sdk, err := fabsdk.New(config.FromFile(ccp))
	if err != nil {
		fmt.Printf("failed to create SDK: %v\n", err)
		return err
	}

	mspClient, err := msp.New(sdk.Context(), msp.WithOrg("Org1")) // TODO get org from somewhere from ccp
	if err != nil {
		fmt.Println("failed to create MSP client")
		return err
	}
	fmt.Println("MSP client created")

	creds, _ := networkOptions.Wallet.Export(networkOptions.Identity)
	switch v := creds.(type) {
	case *X509Identity:
		network.si, err = mspClient.CreateSigningIdentity(othermsp.WithCert([]byte(v.GetCert())), othermsp.WithPrivateKey([]byte(v.GetKey())))
		//fmt.Println("cert:", v.GetCert(), "key:", v.GetKey())
	}

	//fmt.Println(creds)
	network.id = creds
	network.sdk = sdk

	return nil
}

func (network *Network) Dispose() error {
	return nil
}

func (network *Network) GetLedger(channelName string) (*Ledger, error) {
	channel, err := channel.New(network.sdk.ChannelContext(channelName, fabsdk.WithIdentity(network.si)))
	if err != nil {
		fmt.Println("failed to create channel")
		return nil, err
	}

	return newLedger(channel), nil
}

func NewNetwork() Network {
	return Network{}
}
