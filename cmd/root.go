package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "got",
	Short: "Got is a Toy Version Control System",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello World")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
