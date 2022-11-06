package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mercury-client/app"
	logger2 "mercury-client/plugin/logger"
)

var preCommands []func(cmd *cobra.Command, args []string) error

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use: "serve",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		app.RegisterPlugin(logger2.New(logger))

		for _, pre := range preCommands {
			err := pre(cmd, args)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			fmt.Println(err)
			logger.
				WithOptions(zap.AddStacktrace(zap.LevelEnablerFunc(func(zapcore.Level) bool {
					return false
				}))).
				Fatal("Command error", zap.Error(err))
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) error {

	if err := app.Init(); err != nil {
		return err
	}
	defer app.Finish()

	if err := app.Serve(cmd.Context()); err != nil {
		return err
	}
	return nil
}
