package cmd

import (
    "log"
    "runtime/debug"

    "bitbucket.com/eldius/rcon-client-go/internal/rcon"
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
        c, err := rcon.NewClient("localhost:27015")
        if err != nil {
            log.Println(err)
            log.Panicln(string(debug.Stack()))
        }
        pl, err := c.Login("StrongP@ss")
        if err != nil {
            log.Println(err)
            log.Panicln(string(debug.Stack()))
        }
        log.Printf("login response: %v\n", pl.Body)
        pc, err := c.Command("help")
        if err != nil {
            log.Println(err)
            log.Panicln(string(debug.Stack()))
        }
        log.Printf("help response: %s\n", pc.ResponseBody)
        pe, err := c.Command("echo message")
        if err != nil {
            log.Println(err)
            log.Panicln(string(debug.Stack()))
        }
        log.Printf("echo response: %s\n", pe.ResponseBody)
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
