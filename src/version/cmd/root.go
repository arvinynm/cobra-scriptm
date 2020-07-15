package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmn",
	Short: "Show TNM version",
	Long:  "Show Technical Notification Manager version",
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("tnm v1.0")
	//},
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Show TNM version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tnm v1.0")
	},
}

func Execute() {
	rootCmd.Execute()
}
func init() {
	rootCmd.AddCommand(versionCmd)
}