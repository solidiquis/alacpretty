package main

import (
	"log"
	"math"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
}

func themeShuffler(fileContent *string) (*widgets.List, func(chan string)) {
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
	themesList.BorderStyle.Fg = ui.ColorWhite

	currentTheme := currentTheme(fileContent)
	for index, theme := range themesList.Rows {
		if currentTheme == theme {
			themesList.SelectedRow = index
		}
	}

	setThemeState := func(eChan chan string) {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "q":
				eChan <- e.ID
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

func opacityGaugeAdjuster(fileContent *string) (*widgets.Gauge, func(chan string)) {
	opacityGauge := widgets.NewGauge()
	opacityGauge.Title = "Opacity"
	opacityGauge.SetRect(0, 8, 50, 11)
	opacityGauge.Percent = int(currentOpacity(fileContent) * 100)
	opacityGauge.BarColor = ui.ColorYellow
	opacityGauge.LabelStyle = ui.NewStyle(ui.ColorBlue)
	opacityGauge.BorderStyle.Fg = ui.ColorWhite

	setGaugeState := func(eChan chan string) {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			var newOpacity float64
			var tmp int

			switch e.ID {
			case "<C-c>", "q":
				eChan <- e.ID
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

	rowOne := []ui.Drawable{
		themesList,
	}
	rowTwo := []ui.Drawable{
		opacityGauge,
	}
	widgetGrid := [][]ui.Drawable{
		rowOne,
		rowTwo,
	}

	var currentWidget ui.Drawable = themesList // default widget
	currentWidget.(*widgets.List).BorderStyle.Fg = ui.ColorYellow

	ui.Render(
		themesList,
		opacityGauge,
	)

	uiEventChannel := make(chan string)
	go setThemesListState(uiEventChannel)

	var activeRow int
	var activeWidget ui.Drawable
	for {
		e := <-uiEventChannel
		switch e {
		case "<C-j>", "<Down>":
			activeRow--
			if activeRow < 0 {
				activeRow = len(widgetGrid) - int(math.Abs(float64(activeRow)))
			}
		case "<C-k>", "<Up>":
			activeRow++
			if activeRow > len(widgetGrid)-1 {
				activeRow = 0
			}
		case "<C-c>", "q":
			return
		}
		activeWidget = widgetGrid[activeRow][0] // for now..

		// Why isn't this section of code DRY?
		// Drawables have different underlying types hence
		// why type assertion switch statements are necessary.
		switch cw := currentWidget.(type) {
		case *widgets.List:
			cw.BorderStyle.Fg = ui.ColorWhite
			ui.Render(cw)
		case *widgets.Gauge:
			cw.BorderStyle.Fg = ui.ColorWhite
			ui.Render(cw)
		}

		switch aw := activeWidget.(type) {
		case *widgets.List:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			go setThemesListState(uiEventChannel)
		case *widgets.Gauge:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			go setGaugeState(uiEventChannel)
		}
	}
}
