package main

import (
	"fmt"
	"os"

	"github.com/davidkel/simple-fabric-client-go/pkg/api"
)

func main() {

	// wallet test
	idh := api.NewX509IdentityHandler()
	//	fsw := api.NewFileSystemWallet("dirpath", idh)
	//    fsw.Import("dave", api.NewX509Identity("davecert", "davekey"))  // could I have fsw.newIdentity(...string) and it builds the identity object ?
	// or even fsw.import("dave", ...string) and it just does it based on the handler

	imw := api.NewInMemoryWallet(idh)
	// TODO: Need to pull in valid credentials here
	imw.Import("dave", api.NewX509Identity("----BEGIN CERFITICATE---", "-----BEGIN KEY -----"))

	creds, _ := imw.Export("dave")

	switch v := creds.(type) {
	case *api.X509Identity:
		fmt.Println("cert:", v.GetCert(), "key:", v.GetKey())
	}

	/*
		f := api.NewNetwork()
		f.Init("{}", &api.NetworkOptions{Wallet: imw, Identity: "dave"})
		l, _ := f.GetLedger("test")
		c, _ := l.GetContract("test2", nil)
		fmt.Println(c)
	*/

	var (
		channelName = "davechannel"
		ccID        = "mycc"
		//orgName     = "Org1"
		ccpFile = os.Getenv("HOME") + "/othercode/fabric-samples/first-network/connection-org1.yaml"
		myid    = "dave"
	)

	idhandler := api.NewX509IdentityHandler()
	wallet := api.NewInMemoryWallet(idhandler)
	// TODO: Need to pull in valid credentials here
	wallet.Import(myid, api.NewX509Identity("----BEGIN CERFITICATE---", "-----BEGIN KEY -----"))

	f := api.NewNetwork()
	f.Init(ccpFile, &api.NetworkOptions{Wallet: wallet, Identity: myid})
	l, _ := f.GetLedger(channelName)
	c, _ := l.GetContract(ccID)
	c.SubmitTransaction("query", [][]byte{[]byte("a")})

}
