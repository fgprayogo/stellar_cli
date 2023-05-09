/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"stellar_cli/controller"
	"github.com/spf13/cobra"
)

// StartHTTPServiceCmd represents the StartHTTPService command
var StartHTTPServiceCmd = &cobra.Command{
	Use:   "StartHTTPService",
	Short: "Start the HTTP Service",
	Run: func(cmd *cobra.Command, args []string) {
		controller.Init()

	},
}

func init() {
	rootCmd.AddCommand(StartHTTPServiceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// StartHTTPServiceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// StartHTTPServiceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
