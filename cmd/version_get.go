package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	kea "kea-ctrl/kea"
	"net/http"
)

var versionCmdGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get current cache",
	Long: `This command returns extended information about the Kea version that is running. The returned string is the same as if Kea were run with the -V command-line option.

Supported by: kea-ctrl-agent, kea-dhcp-ddns, kea-dhcp4, kea-dhcp6

Availability: 1.2.0 (built-in)`,

	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println(err)
		}
		service,_ := cmd.Flags().GetString("service")

		v := VersionGet(host, service)
		json, _ := json.Marshal(v)
		fmt.Println(string(json))
	},
}

func init() {
	versionCmd.AddCommand(versionCmdGetCmd)
}

func VersionGet(h,s string) interface{} {
	var arg []string
	arg = append(arg, s)

	query := kea.Command{Command: "version-get", Service: arg }

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(query)

	keaAgent := "http://" + h + ":8080/"

	resp, _ := http.Post(keaAgent, "application/json", b)
	body := kea.VersionResults{}
	json.NewDecoder(resp.Body).Decode(&body)

	return body

}