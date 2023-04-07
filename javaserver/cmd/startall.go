/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"javaserver/java"

	"github.com/spf13/cobra"
)

// startallCmd represents the startall command
var startallCmd = &cobra.Command{
	Use:   "startall",
	Short: "开启所有服务",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("startall called")
		for _, service := range ArgsList {
			fmt.Printf("===============================================================\n")
			status, _ := java.CkeckStatus(service)
			if status {
				fmt.Println("该服务正在运行 请先停止")
				continue
			}
			java.Start(service)
			fmt.Printf("===============================================================\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(startallCmd)
}
