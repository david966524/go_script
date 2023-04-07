/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"javaserver/java"

	"github.com/spf13/cobra"
)

// stopallCmd represents the stopall command
var stopallCmd = &cobra.Command{
	Use:   "stopall",
	Short: "停止所有服务",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stopall called")
		for _, service := range ArgsList {
			fmt.Printf("===============================================================\n")
			fmt.Println(service + "stop")
			result := java.Stop(service)
			fmt.Printf("%v 关闭 %v\n", service, result)
			fmt.Printf("===============================================================\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopallCmd)
}
