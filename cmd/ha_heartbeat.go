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

var heartbeatCmd = &cobra.Command{
	Use:   "heartbeat",
	Short: "Get the heartbeat status for service",
	Long: `This command is sent internally by a Kea partner when operating in High-Availability (HA) mode. 
It retrieves the server’s HA state and clock value.

Supported by: kea-dhcp4, kea-dhcp6

Availability: 1.4.0 (high_availability hook library)`,

	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			fmt.Println(err)
		}
		service, _ := cmd.Flags().GetString("service")

		ha := Heartbeat(host, service)
		json, _ := json.Marshal(ha)
		fmt.Println(string(json))
	},
}

var scopesCmd = &cobra.Command{
	Use:   "scopes",
	Short: "This command modifies the scope that the server is responsible for serving when operating in High Availability (HA) mode.",
	Long: `This command modifies the scope that the server is responsible for serving when operating in High Availability (HA) mode.

Supported by: kea-dhcp4, kea-dhcp6

Availability: 1.4.0 (high_availability hook library)

Description and examples: see ha-scopes command

Command syntax:

{
    "command": "ha-scopes",
    "service": [ <service, typically 'dhcp4' or 'dhcp6'> ],
    "arguments": {
        "scopes": [ "HA_server1", "HA_server2" ]
    }
}

In the example below, the arguments configure the server to handle traffic from both the HA_server1 and HA_server2 scopes.

Response syntax:

{
    "result": <integer>,
    "text": "<string>"
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

		scope := Scopes(host, service)
		json, _ := json.Marshal(scope)
		fmt.Println(string(json))
	},
}

func init() {
	haCmd.AddCommand(heartbeatCmd)
	haCmd.AddCommand(scopesCmd)
}

func Heartbeat(h, s string) interface{} {
	var arg []string
	arg = append(arg, s)

	query := kea.Command{Command: "ha-heartbeat", Service: arg}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(query)

	keaAgent := "http://" + h + ":8080/"

	resp, _ := http.Post(keaAgent, "application/json", b)
	body := kea.Results{}
	json.NewDecoder(resp.Body).Decode(&body)

	return body
}

func Scopes(h, s string) interface{} {
	var arg []string
	arg = append(arg, s)

	query := kea.Command{Command: "ha-scopes", Service: arg}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(query)

	keaAgent := "http://" + h + ":8080/"

	resp, _ := http.Post(keaAgent, "application/json", b)
	body := kea.Results{}
	json.NewDecoder(resp.Body).Decode(&body)

	return body
}
