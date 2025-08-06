package expose

import (
	"fmt"
	"github.com/spf13/cobra"
	"dext/utils"
)

var hosts []string
var tlsverify string
var tlscacert string
var tlskey string

// expose represents the daemon command
var ExposeCmd = &cobra.Command{
	Use:   "expose",
	Short: "Expose docker daemon enpoint",
	Long: `This command can help set up docker daemon endpoint`,
	Run: func(cmd *cobra.Command, args []string) {
		for i, host := range hosts {
			fmt.Printf("Host %d: %s\n", i+1, host)
		}
		utils.DebugPrint(tlsverify)
		utils.DebugPrint(tlscacert)
		utils.DebugPrint(tlskey)
	},
}

func init() {
	fmt.Println("[dext/daemon/expose] init")
	ExposeCmd.Flags().StringSliceVarP(&hosts, "host", "H", []string{}, "Docker daemon socket(s) to bind to")
	ExposeCmd.Flags().StringVarP(&tlsverify, "tlsverify", "", "", "tlsverify key path")
	ExposeCmd.Flags().StringVarP(&tlscacert, "tlscacert", "", "", "tlscacert key path")
	ExposeCmd.Flags().StringVarP(&tlskey, "tlskey", "", "", "tlskey key path")
}
