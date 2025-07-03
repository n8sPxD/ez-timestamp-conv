/*
Copyright © 2025 n8sPxD <noobsoap233@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/n8sPxD/tsconv-util/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timeconv [timestamp]",
	Short: "Simple unix timestamp converter",
	Long:  `Simple unix timestamp converter - converts unix timestamp to human readable date format`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timestamp := args[0]

		result, err := internal.ConvertTimestamp(timestamp)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(result)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
