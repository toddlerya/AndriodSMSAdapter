/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/toddlerya/AndriodSMSAdapter/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		server.Server(port, mode)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().Uint16Var(&port, "port", 30000, "设定服务端口")
	serverCmd.PersistentFlags().StringVar(&mode, "mode", "phone", "运行模式:1. phone: 即为在手机内部运行 2. pc: 在PC上外部运行")
}
