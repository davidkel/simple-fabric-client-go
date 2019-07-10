package api

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

type LedgerInterface interface {
	GetContract(ccid string) (Contract, error)
}

type Ledger struct {
	channel *channel.Client
}

func (l *Ledger) GetContract(ccid string) (*Contract, error) {
	return newContract(ccid, l.channel), nil
}

func newLedger(channel *channel.Client) *Ledger {
	return &Ledger{channel}
}
