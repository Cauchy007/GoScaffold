package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server for your application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server...")
		// Add your server start logic here
	},
}

func init() {
	// Register the serve command with the root command
	rootCmd.AddCommand(serveCmd)
}
