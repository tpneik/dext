/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package check

import (
	"fmt"
	"github.com/spf13/cobra"
	"dext/utils"
	"strings"
)

// daemonCmd represents the daemon command
var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of check command",
	Long: `This command is to check if exist docker on this machine or not`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("[dext/daemon/check]: ")
		if cmdStdout, err := utils.Execute("docker", "info"); err != nil {
			fmt.Println(err)
		} else {
			// fmt.Println(cmdStdout)
			lines := strings.Split(cmdStdout, "\n")
			for _, line := range lines {
				if strings.Contains(line, "Docker Root Dir: "){
					parts := strings.SplitN(line, "Docker Root Dir: ", 2 )
					fmt.Println("This is docker home: " + strings.TrimSpace(parts[1]))
				}
			}
		}
	},
}

func init() {
	fmt.Println("[dext/daemon/check] init")
}
