package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
/*
{ "command": "ha-heartbeat", "service": [ "dhcp4","dhcp6" ] }
*/


var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get current version",
	Long: `version`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}