package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/pretty"
	"gitlab.com/ascenty/nemesis-cli/api"
)

var edge string

var loginCmd = &cobra.Command{
	Use:   "login -e <edge>",
	Short: "Login to Nemesis",
	Long:  `Login to Nemesis`,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL := viper.GetString(fmt.Sprintf("%s.baseURL", edge))
		auth := viper.GetString("auth")

		if baseURL == "" || auth == "" {
			return fmt.Errorf("invalid edge provided")
		}

		authResponse, err := api.Login(baseURL, auth)

		if err != nil {
			return fmt.Errorf("error logging in: %s", err)
		}

		authResponseJSON, err := json.Marshal(authResponse)

		if err != nil {
			return fmt.Errorf("error marshalling auth response: %s", err)
		}

		output := pretty.Pretty(authResponseJSON)

		cmd.Print(string(output))
		return nil
	},
}

func init() {
	loginCmd.Flags().StringVarP(&edge, "edge", "e", "edge", "Edge to login to")
	loginCmd.MarkFlagRequired("edge")
}
