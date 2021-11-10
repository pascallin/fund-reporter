package cmd

import (
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"pascal_lin.github.com/fund-reporter/datasource"
	"pascal_lin.github.com/fund-reporter/tui"
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

		log.WithFields(log.Fields{
			"dataField": args[0],
			"dataType":  args[1],
		}).Info("Inside rootCmd Run")

		switch dataField {
		case "stat":
			var startTime, endTime string
			if dataType == datasource.GDP {
				startTime = "2020-A"
				endTime = fmt.Sprintf("%s-A", time.Now().Format("2006"))
			} else {
				startTime = "2020-01"
				endTime = time.Now().Format("2006-01")
			}
			result, err := datasource.GetEconomicStat(dataType, startTime, endTime)
			if err != nil {
				cmd.PrintErr(err)
			}

			log.WithFields(log.Fields{
				"isShowTui": isShowTui,
			}).Debug("tui args")

			if isShowTui && dataType == datasource.GDP {
				tui.GDPBarChart(result.Data)
			} else {
				log.WithField("result", result).Info("Run succeed")
			}
		case "fund":
			code := args[1]
			params := []string{code}

			log.WithFields(log.Fields{
				"params": params,
			}).Info("Crawl fund codes")

			result, err := datasource.GetFundsData(params)
			if err != nil {
				panic(err)
			}
			for _, x := range result {
				log.WithFields(log.Fields{
					"result": x,
					"code":   x.Code,
				}).Info("Run succeed")
			}
		}
	},
}
