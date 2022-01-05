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

// mockCmd represents the mock command
var mockCmd = &cobra.Command{
	Use:   "mock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mock called")
		server.RunMockServer()
	},
}

func init() {
	rootCmd.AddCommand(mockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mockCmd.PersistentFlags().String("foo", "", "A help for foo")
	mockCmd.PersistentFlags().StringVarP(&config.Cfg.Mock.FilePath, config.MockfilePath, "y", "/tmp/mock.yaml", "mock file path")
	mockCmd.PersistentFlags().StringVarP(&config.Cfg.Mock.ServiceType, config.MockService, "s", "server", "mock service type")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
