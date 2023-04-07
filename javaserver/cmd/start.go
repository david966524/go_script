/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"javaserver/java"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "javaserver start 服务名",
	Long:  `start server`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b := ContainsString(ArgsList, args[0])
		if !b {
			fmt.Println("参数错误")
			return
		}
		status, _ := java.CkeckStatus(args[0])
		if status {
			fmt.Println("该服务正在运行 请先停止")
			return
		}
		java.Start(args[0])
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
