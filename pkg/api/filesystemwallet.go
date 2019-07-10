package api

import "path/filepath"
import "fmt"

/*
type Wallet interface {
	Import(Identity) error
    Export(label string) (Identity, error)
    Delete(label string) error
    Exists(label string) error
    //List
}
*/

type FileSystemWallet struct {
	dirPath   string
	idhandler IdHandler
}

func NewFileSystemWallet(dirPath string, idhandler IdHandler) *FileSystemWallet {
	return &FileSystemWallet{dirPath, idhandler}
}

func (f *FileSystemWallet) Import(label string, id IdentityType) error {
	elements := f.idhandler.GetElements(id)
	fileToUse := filepath.Join(f.dirPath, label)
	fmt.Println(fileToUse)
	for key, value := range elements {
		// handle this
		fmt.Println("key:", key, "value:", value)
	}
	return nil
}

/*
func (f *FileSystemWallet) Export(label) IdentityType {
	return nil
}
*/
