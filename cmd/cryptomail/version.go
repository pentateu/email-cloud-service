package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version info",
	Long:  `Every software has a version. This is Guerrilla's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I dont know which version I am :( sorry :( ")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
