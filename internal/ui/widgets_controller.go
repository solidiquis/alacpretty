package ui

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

func WidgetsController(fileContent *string) {
	defer ui.Close()

	themesList, setThemeState := themeShuffler(fileContent)
	opacityGauge, setOpacityState := opacityAdjuster(fileContent)
	fontSizeAdjuster, setFontSizeState := fontSizeAdjuster(fileContent)
	fontShuffler, setFontState := fontShuffler(fileContent)
	themeSearch, setThemeSearchState := themeSearchbox(themesList)
	fontSearch, setFontSearchState := fontSearchbox()

	help := helpBox()
	yaml, yamlChan, setYamlBoxState := showYaml()
	defer close(yamlChan)

	// background
	go setYamlBoxState()

	// default widget
	var currentWidget ui.Drawable = themeSearch
	currentWidget.(*widgets.Paragraph).BorderStyle.Fg = ui.ColorYellow
	yamlChan <- currentWidget.(*widgets.Paragraph).Title

	ui.Render(
		themeSearch,
		themesList,
		opacityGauge,
		fontSizeAdjuster,
		fontShuffler,
		fontSearch,
		yaml,
		help,
	)

	rowOne := []ui.Drawable{
		themeSearch,
		opacityGauge,
	}
	rowTwo := []ui.Drawable{
		themesList,
	}
	rowThree := []ui.Drawable{
		fontSearch,
	}
	rowFour := []ui.Drawable{
		fontShuffler,
		fontSizeAdjuster,
	}
	rowFive := []ui.Drawable{
		help,
	}
	widgetGrid := [][]ui.Drawable{
		rowOne,
		rowTwo,
		rowThree,
		rowFour,
		rowFive,
	}

	var activeRowIndex, activeColumnIndex int
	var activeWidget ui.Drawable

	e := setThemeSearchState() // default active widget

	for {
		switch e {
		case "K":
			activeRowIndex--
		case "J":
			activeRowIndex++
		case "H":
			activeColumnIndex--
		case "L":
			activeColumnIndex++
		case "<C-c>", "q":
			return
		}

		if activeRowIndex < 0 {
			absActiveRowIndex := int(math.Abs(float64(activeRowIndex)))
			activeRowIndex = len(widgetGrid) - absActiveRowIndex
		} else if activeRowIndex > len(widgetGrid)-1 {
			activeRowIndex = 0
		}
		activeRow := widgetGrid[activeRowIndex]

		if activeColumnIndex < 0 {
			absActiveColumnIndex := int(math.Abs(float64(activeColumnIndex)))
			activeColumnIndex = len(activeRow) - absActiveColumnIndex
		} else if activeColumnIndex >= len(activeRow) {
			activeColumnIndex = 0
		}
		activeColumn := activeRow[activeColumnIndex]

		activeWidget = activeColumn

		// yellow = active widget
		// white = inactive widget
		// Deactivate current widget
		switch cw := currentWidget.(type) {
		case *widgets.List:
			cw.BorderStyle.Fg = ui.ColorWhite
			ui.Render(cw)
		case *widgets.Gauge:
			cw.BorderStyle.Fg = ui.ColorWhite
			ui.Render(cw)
		case *widgets.Paragraph:
			cw.BorderStyle.Fg = ui.ColorWhite
			ui.Render(cw)
		}

		// Set new active widget
		switch aw := activeWidget.(type) {
		case *widgets.List:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			yamlChan <- aw.Title

			switch aw.Title {
			case "Themes":
				e = setThemeState()
			case "Size":
				e = setFontSizeState()
			case "Fonts":
				e = setFontState()
			}
		case *widgets.Gauge:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			yamlChan <- aw.Title

			switch aw.Title {
			case "Opacity":
				e = setOpacityState()
			}
		case *widgets.Paragraph:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			yamlChan <- aw.Title

			switch aw.Title {
			case "Search Theme":
				e = setThemeSearchState()
			case "Search Font":
				e = setFontSearchState()
			}
		}
	}
}
