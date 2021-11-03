package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"pascal_lin.github.com/fund-reporter/datasource"
)

var crawlCmd = &cobra.Command{
	Use:   "crawl [stat|fund] [code|gdp|cpi|ppi]",
	Short: "抓取数据",
	Long:  `抓取一些基金数据并展示`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("需要数据类型和具体数据域参数")
		}
		if args[0] != "stat" && args[0] != "fund" {
			return errors.New("需要数据类型只能为[stat|fund]")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataField := args[0]
		dataType := args[1]
		fmt.Printf("Inside rootCmd Run with args: %v\n", args)
		switch dataField {
		case "stat":
			result, err := datasource.GetEconomicStat(dataType)
			if err != nil {
				cmd.PrintErr(err)
			}
			fmt.Printf("result: %v", result)
		}
	},
}
