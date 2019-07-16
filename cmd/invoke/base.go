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

	testPrivKey := `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgzXCAxABlCcgSKC7K
2FD85gNlKXQd3B07O6iRWOMNQBmhRANCAAQgKqnaZhg3b/PdjPsD6KF2b4FeOymm
zkbfZ2h8j6UoUWTi7j72EXNxAvuTwsFQhMjmVWeSJM7UeoF+6qDqfoOW
-----END PRIVATE KEY-----`

	testCert := `-----BEGIN CERTIFICATE-----
MIICKjCCAdGgAwIBAgIRAIBcN+L/lYvyc3HdTF90JW0wCgYIKoZIzj0EAwIwczEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh
Lm9yZzEuZXhhbXBsZS5jb20wHhcNMTkwNzA4MTA1NjAwWhcNMjkwNzA1MTA1NjAw
WjBsMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN
U2FuIEZyYW5jaXNjbzEPMA0GA1UECxMGY2xpZW50MR8wHQYDVQQDDBZVc2VyMUBv
cmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEICqp2mYY
N2/z3Yz7A+ihdm+BXjspps5G32dofI+lKFFk4u4+9hFzcQL7k8LBUITI5lVnkiTO
1HqBfuqg6n6DlqNNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwKwYD
VR0jBCQwIoAgNnJQPzBPUFVRASHxzAnuTaaLm37Uh0EGOyqAd4j7SXcwCgYIKoZI
zj0EAwIDRwAwRAIgQUjPUIJxA+n7fXe4CL++Q1PaBELBW3EjgrXmpiuUbUoCIH7I
EHDlZdXELzv1alY4GavheQ+L9rMDKlnw0/N4HDTT
-----END CERTIFICATE-----`

	idhandler := api.NewX509IdentityHandler()
	wallet := api.NewInMemoryWallet(idhandler)
	wallet.Import(myid, api.NewX509Identity(testCert, testPrivKey))

	f := api.NewNetwork()
	f.Init(ccpFile, &api.NetworkOptions{Wallet: wallet, Identity: myid})
	l, _ := f.GetLedger(channelName)
	c, _ := l.GetContract(ccID)
	r, _ := c.Query("query", [][]byte{[]byte("a")})
	fmt.Println(string(r))

}
