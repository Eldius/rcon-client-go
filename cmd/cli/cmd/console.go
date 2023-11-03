package cmd

import (
	"bitbucket.com/eldius/rcon-client-go/helper"
	"bitbucket.com/eldius/rcon-client-go/internal/config"
	"bitbucket.com/eldius/rcon-client-go/internal/output"
	"bitbucket.com/eldius/rcon-client-go/protocol"
	"fmt"
	"github.com/pterm/pterm/putils"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Opens an RCON interactive console",
	Long:  `Opens an RCON interactive console.`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.DebugMode() {
			pterm.EnableDebugMessages()
		}

		s, err := pterm.DefaultBigText.WithLetters(putils.LettersFromStringWithRGB("rcon console", pterm.NewRGB(0, 9, 0))).Srender()
		if err != nil {
			pterm.Error.Println("Failed to start output:", err)
		}
		pterm.DefaultCenter.Println(pterm.Green(s))
		c, err := protocol.NewClient(protocol.WithHost(consoleHost), protocol.WithDebugLog(config.DebugMode()), protocol.WithWriter(&output.MyWriter{}))
		if err != nil {
			pterm.Error.Println("Failed to connect:", err)
			os.Exit(1)
		}
		defer func(c *protocol.Client) {
			err := c.Close()
			if err != nil {
				pterm.Error.Println("Failed to disconnect from server:", err)
			}
		}(c)
		pass, err := helper.AskForPassword("server password:")
		if err != nil {
			pterm.Error.Println("Failed to get password:", err)
			os.Exit(1)
		}
		_, err = c.Login(pass)
		if err != nil {
			pterm.Error.Println("Failed to log in:", err)
			os.Exit(1)
		}
		pterm.DefaultHeader.FullWidth = true
		pterm.DefaultHeader.Println("Connected")
		for {
			command, err := readCommand("$: ")
			if err != nil {
				pterm.Error.Println("Failed to read command:", err)
				os.Exit(1)
			}
			consoleDebug(
				fmt.Sprintf("cmd as byte:   -->%v<--", []byte(command)),
				fmt.Sprintf("cmd as string: -->%s<--", command))

			if ("exit" == command) || ("quit" == command) {
				pterm.Info.Println("Closing console...")
				os.Exit(0)
			}
			res, err := c.Command(command)
			if err != nil {
				pterm.Error.Println("Failed to execute command:", err)
				os.Exit(1)
			}
			showCommandOutput(res)
		}
	},
}

func showCommandOutput(p *protocol.Packet) {
	pterm.DefaultHeader.Println("Execution result")
	pterm.DefaultBasicText.Printfln("id:       %d => %d", p.ID, p.ResponseID)
	pterm.DefaultBasicText.Printfln("type:     %s => %s", p.Type, p.ResponseType)
	pterm.DefaultBasicText.Printfln("cmd:      %s", p.Body)
	pterm.DefaultBasicText.Printfln("response: %s", p.ResponseBody)

}

func consoleDebug(msgs ...string) {
	if config.DebugMode() {
		pterm.Debug.Println(pterm.Blue("[console] -- console debug -----"))
		for _, msg := range msgs {
			pterm.Debug.Printfln(pterm.Blue(pterm.Sprintf("[console] %s", msg)))
		}
		pterm.Debug.Println(pterm.Blue("[console] --------------------"))
	}
}

func readCommand(prompt string) (string, error) {
	command, err := pterm.DefaultInteractiveTextInput.WithMultiLine(false).Show(prompt)
	if err != nil {
		pterm.Error.Println("Failed to read command:", err)
		os.Exit(1)
	}

	pterm.Println() // Blank line
	command = strings.Trim(command, "\n")
	command = strings.Trim(command, "\r") // to run in Windows OS
	return command, err
}

var (
	consoleHost string
)

func init() {
	rootCmd.AddCommand(consoleCmd)
	consoleCmd.Flags().StringVarP(&consoleHost, "server", "s", "localhost:27015", "Remote server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consoleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consoleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
