/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"javaserver/java"

	"github.com/spf13/cobra"
)

// rebootallCmd represents the rebootall command
var restartallCmd = &cobra.Command{
	Use:   "restartall",
	Short: "重启所有服务",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restartall called")
		for _, service := range ArgsList {
			fmt.Printf("===============================================================\n")
			fmt.Printf("%v restart \n", service)
			java.Restart(service)
			fmt.Printf("===============================================================\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(restartallCmd)

}
