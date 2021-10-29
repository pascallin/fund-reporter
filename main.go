package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"pascal_lin.github.com/spider/datasource"
)

func main() {
	codes := []string{"161725", "481010"}

	fName := fmt.Sprintf("dataset_%s.csv", time.Now().Format("20211029"))
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Name", "Code", "CurrentPrice"}
	writer.Write(header)

	result := datasource.GetFundsData(codes)
	// result := datasource.GetFundsDataWithQueue(2)
	for _, x := range result {
		fmt.Printf("%v\n", x)
		writer.Write([]string{x.Name, x.Code, strconv.FormatFloat(x.CurrentPrice, 'f', 4, 64)})
	}
}
