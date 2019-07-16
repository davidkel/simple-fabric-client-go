package api

import "testing"
import mock "github.com/davidkel/simple-fabric-client-go/pkg/api/mocks/channel"

func TestSubmitTransaction(t *testing.T) {

	var x *mock.Client
	x = &mock.Client{}
	c := &Contract{"someid", x}
	c.SubmitTransaction("fgfg", [][]byte{})

}
