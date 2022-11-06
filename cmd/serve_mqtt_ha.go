//go:build ha

package cmd

import (
	"github.com/spf13/cobra"
	"mercury-client/app"
	"mercury-client/plugin/ha"
)

func init() {
	const flagPrefix = "ha"
	serveCmd.Flags().String(flagPrefix, "homeassistant", "MQTT prefix for HA integration")
	preCommands = append(preCommands, func(cmd *cobra.Command, args []string) error {
		prefix, err := serveCmd.Flags().GetString(flagPrefix)
		if err != nil {
			return err
		}
		app.RegisterPlugin(ha.New(mqttPlugin, prefix))
		return nil
	})
}
