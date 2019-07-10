package api

type LedgerInterface interface {
	GetContract(ccid string) (Contract, error)
}

type Ledger struct {
}

func (l *Ledger) GetContract(ccid string) (*Contract, error) {
	return newContract("fred"), nil
}

func newLedger() *Ledger {
    return &Ledger{}
}
