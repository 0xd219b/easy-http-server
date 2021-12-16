/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/Berger7/easy-http-server/config"
	"github.com/Berger7/easy-http-server/server"
	"github.com/spf13/cobra"
)

// singleCmd represents the single command
var singleCmd = &cobra.Command{
	Use:   "single",
	Short: "A single server(POST and GET only)",
	Long: `A single server, use this command to run a single server:

Only support GET and POST request method.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("single called")
		server.RunSingleServer()
	},
}

func init() {
	rootCmd.AddCommand(singleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// singleCmd.PersistentFlags().String("foo", "", "A help for foo")
	singleCmd.PersistentFlags().StringVarP(&config.Cfg.RequestMethod, config.RequestMethod, "m", "GET", "request method")
	singleCmd.PersistentFlags().StringVarP(&config.Cfg.RequestUrl, config.RequestUrl, "u", "/", "request url")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// singleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
