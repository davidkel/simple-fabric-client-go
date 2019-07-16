package channel

import sdkchannel "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
import "fmt"

type ClientI interface {
	Execute(r sdkchannel.Request, o ...sdkchannel.RequestOption) (sdkchannel.Response, error)
	Query(r sdkchannel.Request, o ...sdkchannel.RequestOption) (sdkchannel.Response, error)
}

type Client struct{}

func (m *Client) Execute(r sdkchannel.Request, o ...sdkchannel.RequestOption) (sdkchannel.Response, error) {
	fmt.Println("Execute called")
	return sdkchannel.Response{}, nil
}

func (m *Client) Query(r sdkchannel.Request, o ...sdkchannel.RequestOption) (sdkchannel.Response, error) {
	fmt.Println("Query Called")
	return sdkchannel.Response{}, nil
}
