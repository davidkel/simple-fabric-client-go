package main

import "fmt"
import "github.com/davidkel/simple-fabric-client-go/pkg/api"

func main() {

	// wallet test
	idh := api.NewX509IdentityHandler()
	//	fsw := api.NewFileSystemWallet("dirpath", idh)
	//    fsw.Import("dave", api.NewX509Identity("davecert", "davekey"))  // could I have fsw.newIdentity(...string) and it builds the identity object ?
	// or even fsw.import("dave", ...string) and it just does it based on the handler

	imw := api.NewInMemoryWallet(idh)
	imw.Import("dave", api.NewX509Identity("----BEGIN CERFITICATE---", "-----BEGIN KEY -----"))

	creds, _ := imw.Export("dave")

	switch v := creds.(type) {
	case *api.X509Identity:
		fmt.Println("cert:", v.GetCert(), "key:", v.GetKey())
	}

	f := api.NewNetwork()
	f.Init("{}", &api.NetworkOptions{Wallet: imw, Identity: "dave"})
	l, _ := f.GetLedger("test")
	c, _ := l.GetContract("test2")
	fmt.Println(c)

}
