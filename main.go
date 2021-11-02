package main

import "pascal_lin.github.com/spider/datasource"

func main() {
	// datasource.DownloadFundData()
	datasource.GetEconomicStat(datasource.PPI)
}
