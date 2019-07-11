package api

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
)

type ContractInterface interface {
	SubmitTransaction() error
	Query() (string, error)
}

type Contract struct {
	ccid    string
	channel *channel.Client
}

func (c *Contract) SubmitTransaction(fcn string, args [][]byte) error {
	fmt.Println("About to execute")
	_, err := c.channel.Execute(channel.Request{ChaincodeID: c.ccid, Fcn: fcn, Args: args},
		channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		fmt.Printf("Guess it didn't work: %s\n", err)
	}
	fmt.Println("Guess it worked")
	return nil
}
func (c *Contract) Query(fcn string, args [][]byte) ([]byte, error) {
	response, err := c.channel.Query(channel.Request{ChaincodeID: c.ccid, Fcn: fcn, Args: args})
	if err != nil {
		fmt.Printf("failed to query chaincode: %s\n", err)
	}
	return response.Payload, nil
}

func newContract(ccid string, channel *channel.Client) *Contract { //TODO should strings be pointers ?
	return &Contract{ccid, channel}
}
