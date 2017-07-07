package astilectron_bindata

import "github.com/asticode/go-astilectron"

const (
	SrcFileNameAstilectron = "astilectron.zip"
	SrcFileNameElectron    = "electron.zip"
)

func NewProvisioner(disembedFunc func(string) ([]byte, error)) astilectron.Provisioner {
	return astilectron.NewDisembedderProvisioner(disembedFunc, SrcFileNameAstilectron, SrcFileNameElectron)
}
