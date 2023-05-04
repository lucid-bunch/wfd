package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wfd",
	Short: "wfd is a tool that helps with dinner planning",
	Long:  "What's For Dinner is a tool that helps with dinner planning",
}

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(weekCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
