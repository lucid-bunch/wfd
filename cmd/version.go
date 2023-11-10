package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var build = "latest"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wfd",
	Long:  `All software has versions. This is wfd's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("wfd - What's for dinner [%s]\n", build)
	},
}
