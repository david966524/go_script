/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"javaserver/java"

	"github.com/spf13/cobra"
)

var ArgsList []string = []string{"web", "backend", "jobs", "games", "im", "gateway", "eureka", "xxl-job"}

// rootCmd represents the base command when called without any subcommands
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "javaserver stop 服务名",
	Long: `stop server
	web, backend, jobs, games, im, gateway, eureka, xxl-job`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		b := ContainsString(ArgsList, args[0])
		if !b {
			fmt.Println("参数错误")
			return
		}

		fmt.Printf("stop server\n")
		result := java.Stop(args[0])
		fmt.Printf("%v 关闭 %v \n", args[0], result)
	},
}

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
