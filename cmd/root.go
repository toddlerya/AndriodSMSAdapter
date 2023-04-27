/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/toddlerya/glue/logger"
)

var (
	logLevel string
	logJSON  bool
	port     uint16
	mode string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "AndriodSMSAdapter",
	Short: "获取安卓短信的API服务",

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// 初始化日志配置
		if logJSON {
			// 初始化日志配置
			logger.InitLogConfig("json", "./", "android_sms_adapter.log")
		} else {
			logger.InitLogConfig("text", "./", "android_sms_adapter.log")
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		logger.SetLogLevel(logLevel)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "INFO", "设置日志级别")
	rootCmd.PersistentFlags().BoolVar(&logJSON, "log-json", false, "设置日志为JSON格式")
	err := rootCmd.Execute()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.AndriodSMSAdapter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
