/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package daemon

import (
	"fmt"
	"github.com/spf13/cobra"
	"dext/cmd/daemon/check"
	"dext/cmd/daemon/expose"
)

// daemonCmd represents the daemon command
var DaemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[dext/daemon] AAAAAAAAAaaaaa Daemon neeee")
	},
}

func init() {
	// cmd.rootCmd.AddCommand(daemonCmd)
	fmt.Println("[dext/daemon] init")
	DaemonCmd.AddCommand(check.CheckCmd)
	DaemonCmd.AddCommand(expose.ExposeCmd)
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
