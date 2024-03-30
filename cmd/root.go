/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/8VIM/keyboard_layout_calculator/cli"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var cliConfig *cli.Config = &cli.Config{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "keyboard_layout_calculator",
	Run: func(cmd *cobra.Command, args []string) { cli.Execute(cliConfig, args) },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/keyboard_layout_calculator/config.yaml)")
	rootCmd.PersistentFlags().CountVarP(&cliConfig.Verbose, "verbose", "v", "log verbosity")
	rootCmd.PersistentFlags().BoolVarP(&cliConfig.Force, "force", "f", false, "overwrite cache")
	rootCmd.PersistentFlags().IntVarP(&cliConfig.Parallelism, "parallelism", "p", 2, "Download parallelism")
	rootCmd.PersistentFlags().StringVarP(&cliConfig.Output, "output", "o", ".", "Output directory")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".a" (without extension).
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("$HOME/.config/keyboard_layout_calculator")
		viper.AddConfigPath(".")
	}
	err := viper.ReadInConfig()
	cobra.CheckErr(err)

	err = viper.Unmarshal(&cliConfig.Config, func(c *mapstructure.DecoderConfig) {
		c.ErrorUnused = true
		c.ErrorUnset = true
	})
	cobra.CheckErr(err)
}
