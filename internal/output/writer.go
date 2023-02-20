package output

import (
	"bitbucket.com/eldius/rcon-client-go/protocol"
	"github.com/pterm/pterm"
)

type MyWriter struct {
	protocol.Writer
}

func (w *MyWriter) Write(msgs ...string) {
	pterm.Println()
	pterm.Debug.Println("-- debug -----")
	for _, msg := range msgs {
		pterm.Debug.Printfln("%s", msg)
	}
	pterm.Println()
}
