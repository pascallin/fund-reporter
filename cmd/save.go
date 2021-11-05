package cmd

import (
	"github.com/spf13/cobra"
	"pascal_lin.github.com/fund-reporter/datasource"
)

var saveCmd = &cobra.Command{
	Use:   "save code1 code2 ...",
	Short: "基金数据另存为",
	Long:  `将基金抓取的数据另存为csv文件`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		datasource.DownloadFundData(args[1:], saveFileName)
	},
}
