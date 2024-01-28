package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Long:  `Start the server for your application`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server:[pre-run]...")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the server:[run]...")
		// Add your server start logic here
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(context.Background(), err)
	}
	return nil
}
