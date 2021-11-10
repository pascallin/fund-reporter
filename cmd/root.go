package cmd

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "fund-reporter",
		Short: "一个中国基金及经济数据命令行工具",
		Long: `获取中国基金及经济数据的命令行工具。
致在收集数据，可视化，并以文件格式进行导出方便做其他数据分析。`,
	}

	isShowTui    bool
	saveFileName string
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// logger config
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	// init env
	cobra.OnInitialize(initConfig)

	// crawl command flags
	crawlCmd.PersistentFlags().BoolVarP(&isShowTui, "tui", "t", false, "用terminal UI展示数据")

	// save command flags
	defaultFileName := fmt.Sprintf("dataset_%s.csv", time.Now().Format("2006-01-02"))
	saveCmd.PersistentFlags().StringVarP(&saveFileName, "file", "f", defaultFileName, "保存文件名")

	// disable completion command temporary
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.AddCommand(saveCmd)
	rootCmd.AddCommand(crawlCmd)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
