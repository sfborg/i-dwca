package idwca

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Here is a real implementation of Greeter
type ImporterCol struct {
	Logger hclog.Logger
}

func (g *ImporterCol) Import() error {
	g.Logger.Debug("ImporterCol Imports Stuff")
	fmt.Print("IMPORTTTT")
	f, err := os.Create("FILE")
	if err != nil {
		return err
	}
	defer f.Close()

	return nil
}

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "IMPORT_PLUGIN",
	MagicCookieValue: "import",
}
