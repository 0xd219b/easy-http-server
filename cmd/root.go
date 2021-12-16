/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/Berger7/easy-http-server/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "easy-http-server",
	Short: "An easy way to start a server for testing",
	Long: `The easy-http-server can help you start http server by command with
an easy way. For example:

# 👇 start a single server with POST method and url /hooks
./easy-http-server single -u hooks -m POST

More powerful commands are coming soon.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.easy-http-server.yaml)")
	singleCmd.PersistentFlags().Int64VarP(&config.Cfg.ServerPort, config.ServerPort, "p", 8080, "server port")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
