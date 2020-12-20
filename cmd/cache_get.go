package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	kea "kea-ctrl/kea"
	"net/http"
)

/*
{ "command": "ha-heartbeat", "service": [ "dhcp4","dhcp6" ] }
*/

var cacheCmdGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get current cache",
	Long: `This command returns the full content of the host cache.

Supported by: kea-dhcp4, kea-dhcp6

Availability: 1.4.0 (host_cache hook library)


Response syntax:

{
    "result": 0,
    "text": "123 entries returned.",
    "arguments": <list of host reservations>
}

Result is an integer representation of the status. Currently supported statuses are:

    0 - success
    1 - error
    2 - unsupported
    3 - empty (command was completed successfully, but no data was affected or returned)
`,
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println(err)
		}
		service, _ := cmd.Flags().GetString("service")

		fmt.Println(Get(host, service))
	},
}

func init() {
	cacheCmd.AddCommand(cacheCmdGetCmd)
}

func Get(h, s string) interface{} {
	var arg []string
	arg = append(arg, s)

	query := kea.Command{Command: "cache-get", Service: arg}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(query)

	keaAgent := "http://" + h + ":8080/"

	resp, _ := http.Post(keaAgent, "application/json", b)
	body := kea.Results{}
	json.NewDecoder(resp.Body).Decode(&body)

	return body
}
