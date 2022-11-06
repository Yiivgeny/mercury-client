//go:build mqtt || ha

package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"mercury-client/app"
	"mercury-client/plugin/mqtt"
	"net/url"
)

var mqttPlugin *mqtt.Plugin

func init() {
	const flagName = "mqtt"

	serveCmd.Flags().StringArray(flagName, nil, "MQTT broker connection string")
	cobra.CheckErr(viper.BindPFlag(flagName, serveCmd.Flags().Lookup(flagName)))

	mqttEnv := viper.New()
	mqttEnv.SetDefault("protocol", "tcp")
	mqttEnv.SetDefault("port", "1883")

	mqttEnv.SetEnvPrefix("mqtt")
	mqttEnv.MustBindEnv("host")
	mqttEnv.MustBindEnv("port")
	mqttEnv.MustBindEnv("protocol")
	mqttEnv.MustBindEnv("username")
	mqttEnv.MustBindEnv("password")

	preCommands = append(preCommands, func(cmd *cobra.Command, args []string) error {
		brokers := viper.GetStringSlice(flagName)
		if len(brokers) == 0 && mqttEnv.IsSet("host") {
			brokers = append(brokers, fmt.Sprintf("%s://%s:%s",
				mqttEnv.Get("protocol"),
				mqttEnv.Get("host"),
				mqttEnv.Get("port"),
			))
		}
		for i, broker := range brokers {
			parsed, err := url.Parse(broker)
			if err != nil {
				return errors.WithMessage(err, "parsing mqtt broker url")
			}
			if parsed.User == nil && mqttEnv.IsSet("username") {
				parsed.User = url.User(mqttEnv.GetString("username"))
			}
			if parsed.User != nil && mqttEnv.IsSet("password") {
				_, passwordSet := parsed.User.Password()
				if !passwordSet {
					parsed.User = url.UserPassword(parsed.User.Username(), mqttEnv.GetString("password"))
				}
			}
			brokers[i] = parsed.String()
		}

		mqttPlugin = mqtt.New(
			logger.
				//With(zap.String("component", "mqtt")).
				WithOptions(
					zap.AddStacktrace(mqtt.LevelEnablerNone),
					zap.IncreaseLevel(zap.InfoLevel),
				),
			brokers,
		)
		app.RegisterPlugin(mqttPlugin)
		return nil
	})
}
