package api

type ContractInterface interface {
	SubmitTransaction() error
	Query() (string, error)
}

type Contract struct {
    info string
}

func (c *Contract) SubmitTransaction() error {
	return nil
}
func (c *Contract) Query() (string, error) {
	return "", nil
}

func newContract(info string) *Contract {
    return &Contract{info}
}