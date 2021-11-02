package main

import (
	"fmt"

	"pascal_lin.github.com/spider/datasource"
)

func main() {
	// datasource.DownloadFundData()
	result, err := datasource.GetEconomicStat(datasource.PPI)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Results: %v\n", result)
}
