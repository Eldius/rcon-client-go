package cmd

import (
	"bitbucket.com/eldius/rcon-client-go/internal/helper"
	"bitbucket.com/eldius/rcon-client-go/internal/rcon"
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Opens an RCON console",
	Long:  `Opens an RCON console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("console called")
		c, err := rcon.NewClient(consoleHost)
		if err != nil {
			fmt.Println("Failed to connect:", err)
			os.Exit(1)
		}
		pass, err := helper.AskForPassword("server password:")
		if err != nil {
			fmt.Println("Failed to get password:", err)
			os.Exit(1)
		}
		_, err = c.Login(pass)
		if err != nil {
			fmt.Println("Failed to log in:", err)
			os.Exit(1)
		}
		fmt.Println("Logged in  successfully.")
		for {
			fmt.Print("$: ")
			reader := bufio.NewReader(os.Stdin)
			command, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Failed to read command:", err)
				os.Exit(1)
			}
			if "exit" == command || "quit" == command {
				fmt.Println("Closing console...")
				os.Exit(0)
			}
			res, err := c.Command(command)
			if err != nil {
				fmt.Println("Failed to execute command:", err)
				os.Exit(1)
			}
			fmt.Printf("%s\n", res.ResponseBody)
		}
	},
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
