package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
}

func themeShuffler(fileContent *string) (*widgets.List, func()) {
	var rows []string
	for theme := range allThemes {
		rows = append(rows, theme)
	}

	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = rows
	themesList.TextStyle = ui.NewStyle(ui.ColorYellow)
	themesList.WrapText = false
	themesList.SetRect(0, 0, 25, 8)

	currentTheme := currentTheme(fileContent)
	for index, theme := range themesList.Rows {
		if currentTheme == theme {
			themesList.SelectedRow = index
		}
	}

	setThemeState := func() {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>", "<Enter>":
				return
			case "j", "<Down>":
				themesList.ScrollDown()
			case "k", "<Up>":
				themesList.ScrollUp()
			}

			currentTheme := themesList.Rows[themesList.SelectedRow]
			changeTheme(fileContent, currentTheme)
			applyChanges(*fileContent)

			ui.Render(themesList)
		}
	}

	return themesList, setThemeState
}

func opacityGaugeAdjuster(fileContent *string) (*widgets.Gauge, func()) {
	opacityGauge := widgets.NewGauge()
	opacityGauge.Title = "Opacity"
	opacityGauge.SetRect(0, 8, 50, 11)
	opacityGauge.Percent = int(currentOpacity(fileContent) * 100)
	opacityGauge.BarColor = ui.ColorYellow
	opacityGauge.LabelStyle = ui.NewStyle(ui.ColorBlue)
	opacityGauge.BorderStyle.Fg = ui.ColorWhite

	setGaugeState := func() {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			var newOpacity float64
			var tmp int

			switch e.ID {
			case "q", "<C-c>", "<Enter>":
				return
			case "l":
				tmp = opacityGauge.Percent + 10
			case "h":
				tmp = opacityGauge.Percent - 10
			default:
				tmp = 1925 // Cthulhu Fhtagn
			}

			if tmp >= 0 && tmp <= 100 {
				newOpacity = float64(tmp) / 100
				opacityGauge.Percent = tmp
				changeOpacity(fileContent, newOpacity)
				applyChanges(*fileContent)
				ui.Render(opacityGauge)
			}
		}
	}

	return opacityGauge, setGaugeState
}

func widgetsController(fileContent *string) {
	defer ui.Close()
	themesList, setThemesListState := themeShuffler(fileContent)
	opacityGauge, setGaugeState := opacityGaugeAdjuster(fileContent)

	widgets := []ui.Drawable{
		themesList,
		opacityGauge,
	}

	ui.Render(widgets...)

	//uiEvents := ui.PollEvents()

	// theme shuffler is the default active widget
	//setThemesListState()

	// Bookmark
	//var i int
	//for {
	//e := <-uiEvents
	//switch e.ID {
	//case "<C-j>", "<Down>":

	//}
	//}

	setGaugeState()
	setThemesListState()
}
