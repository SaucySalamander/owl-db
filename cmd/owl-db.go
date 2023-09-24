package cmd

import (
	"github.com/SaucySalamander/owl-db/internal/config"
	"github.com/spf13/cobra"
)

var configName string

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.SetupEnv(configName); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVar(&configName, "config", "", "name of the config to use")
}

func Execute() error {
	return rootCmd.Execute()
}
