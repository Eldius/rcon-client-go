package cmd

import (
	"bitbucket.com/eldius/rcon-client-go/internal/config"
	"bitbucket.com/eldius/rcon-client-go/protocol"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//c, err := protocol.NewClient("192.168.0.12:25575")
		c, err := protocol.NewClient(protocol.WithHost("127.0.0.1:27015"), protocol.WithDebugLog(config.DebugMode()))
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			os.Exit(1)
		}
		defer func() {
			_ = c.Close()
		}()

		fmt.Println("Starting login process...")
		pl, err := c.Login("StrongP@ss")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			fmt.Printf("[error] login response: (id: %d => %d) '%s'\n", pl.ID, pl.ResponseID, pl.ResponseBody)
			os.Exit(1)
		}
		fmt.Println("Login success!")
		fmt.Printf("login response: (id: %d => %d) '%s'\n", pl.ID, pl.ResponseID, pl.ResponseBody)
		fmt.Println("-------------")

		fmt.Println("Starting help command...")
		pc, err := c.Command("help")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			fmt.Printf("[error] help response: (id: %d => %d)\n%s\n", pc.ID, pc.ResponseID, pc.ResponseBody)
			os.Exit(1)
		}
		fmt.Printf("help response: (id: %d => %d)\n%s\n", pc.ID, pc.ResponseID, pc.ResponseBody)
		fmt.Println("-------------")
		fmt.Println("Starting new help command...")
		ph, err := c.Command("help")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			fmt.Printf("[error] new help response: (id: %d => %d)\n%s\n", ph.ID, ph.ResponseID, ph.ResponseBody)
			os.Exit(1)
		}
		fmt.Printf("help new response: (id: %d => %d)\n%s\n", ph.ID, ph.ResponseID, ph.ResponseBody)
		fmt.Println("-------------")

		fmt.Println("Starting message command...")
		pm, err := c.Command("say Server is restarting!")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			fmt.Printf("[error] message response:  (id: %d => %d)\n'%s'\n", pm.ID, pm.ResponseID, pm.ResponseBody)
			os.Exit(1)
		}
		fmt.Printf("message response:  (id: %d => %d)\n'%s'\n", pm.ID, pm.ResponseID, pm.ResponseBody)
		fmt.Println("-------------")

		fmt.Println("Starting seed command...")
		ps, err := c.Command("seed")
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack()))
			fmt.Printf("[error] seed response:  (id: %d => %d)\n'%s'\n", ps.ID, ps.ResponseID, ps.ResponseBody)
			os.Exit(1)
		}
		fmt.Printf("seed response:  (id: %d => %d)\n'%s'\n", ps.ID, ps.ResponseID, ps.ResponseBody)
		fmt.Println("-------------")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
