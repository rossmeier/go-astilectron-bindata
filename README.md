# go-astilectron-bindata
go-astilectron-bindata is a simple utility, that automatically downloads [go-astilectron](https://github.com/asticode/go-astilectron)'s
binary dependencies and embeds them into the source code.
## Installation
`go get github.com/veecue/go-astilectron-bindata/...`

## Usage
 * add this line to your code: `//go:generate go-astilectron-bindata` and run `go generate`
   *or* simply run `go-astilectron-bindata`
 * add `"github.com/veecue/go-astilectron-bindata"` to your includes
 * add this code before `a.Start()`: `a.SetProvisioner(astilectron_bindata.NewProvisioner(Disembed))`
