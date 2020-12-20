package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// haCmd represents the ha command
var haCmd = &cobra.Command{
	Use:   "ha",
	Short: "High Availability plugin",
	Long: `Kea's high availability plugin commands

used for DHCP4 and DHCP6 services'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ha called")
	},
}

func init() {
	rootCmd.AddCommand(haCmd)
	//rootCmd.PersistentFlags().StringP("service", "s", "", "Kea service")
	viper.BindPFlag("service", rootCmd.PersistentFlags().Lookup("service"))
}
