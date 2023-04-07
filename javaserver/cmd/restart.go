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
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "javaserver restart 服务名",
	Args:  cobra.MaximumNArgs(1),
	Long:  `restart java service`,
	Run: func(cmd *cobra.Command, args []string) {
		b := ContainsString(ArgsList, args[0])
		if !b {
			fmt.Println("参数错误")
			return
		}
		java.Restart(args[0])
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
