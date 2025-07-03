/*
Copyright © 2025 n8sPxD <noobsoap233@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/n8sPxD/ez-timestamp-conv/internal"

	"github.com/spf13/cobra"
)

var dateFlag string
var millisFlag bool

var rootCmd = &cobra.Command{
	Use:   "timeconv [timestamp]",
	Short: "简单的unix时间戳转换工具",
	Long:  `简单的unix时间戳转换工具 - 支持unix时间戳与人类可读日期格式之间的双向转换`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 检查是否提供了date参数
		if dateFlag != "" {
			// 转换日期为时间戳
			result, err := internal.ConvertDateToTimestamp(dateFlag, millisFlag)
			if err != nil {
				fmt.Printf("错误: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(result)
			return
		}

		// 如果没有date参数，处理时间戳转换
		if len(args) == 0 {
			// 没有参数时显示帮助信息
			cmd.Help()
			return
		}

		timestamp := args[0]
		result, err := internal.ConvertTimestamp(timestamp)
		if err != nil {
			fmt.Printf("错误: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&dateFlag, "date", "d", "", "要转换的日期 (格式: 2006-01-02 15:04:05 或 2006-01-02 15:04)")
	rootCmd.Flags().BoolVarP(&millisFlag, "millis", "m", false, "输出毫秒级时间戳")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
