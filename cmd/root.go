// Package cmd contains all CLI commands used by the application.
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

var rootCmd *cobra.Command

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(NewRootCommand().Execute())
}

// NewRootCommand returns the root command for the CLI.
func NewRootCommand() *cobra.Command {
	rootCmd = &cobra.Command{
		Use:     "picker [command]",
		Version: "0.0.0",
		Long:    "picker returns one or more items from a list",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initConfig()
		},
		Run: func(cmd *cobra.Command, args []string) {
			// does nothing without subcommands
			_ = cmd.Help()
		},
	}

	return rootCmd
}

func initConfig() {
	// set a default logging level to warning at least while we set everything up
	log.SetLevel(log.WarnLevel)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	h, err := os.UserHomeDir()
	if err != nil {
		log.WithError(err).Warn("getting home directory to find config")
	} else {
		viper.AddConfigPath(filepath.Join(h, "/.picker/"))
	}

	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.WithField("file", viper.GetViper().ConfigFileUsed()).
			WithError(err).
			Error("reading config file")

		return
	}

	level, err := log.ParseLevel(viper.GetString("logging.level"))
	if err != nil {
		log.WithError(err).Warn("setting logging level")
	}

	log.SetLevel(level)
}
