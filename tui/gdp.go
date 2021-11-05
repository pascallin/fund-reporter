package tui

import (
	"log"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"pascal_lin.github.com/fund-reporter/datasource"
)

func GDPBarChart(dataset []datasource.ResData) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	bc := widgets.NewBarChart()
	var data []float64
	var labels []string

	for _, item := range dataset {
		pickData := item.Values[len(item.Values)-1]
		floatData, err := strconv.ParseFloat(pickData, 64)
		if err != nil {
			panic(err)
		}
		data = append(data, floatData)
		labels = append(labels, item.Datetime)
	}

	bc.Data = data
	bc.Labels = labels

	bc.Title = "GDP Bar Chart"
	bc.SetRect(5, 5, 120, 25)
	bc.BarWidth = 20
	bc.BarColors = []ui.Color{ui.ColorGreen}
	bc.LabelStyles = []ui.Style{ui.NewStyle(ui.ColorBlue)}
	bc.NumStyles = []ui.Style{ui.NewStyle(ui.ColorYellow)}

	ui.Render(bc)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		}
	}
}
