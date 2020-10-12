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
	opacityGauge, SetOpacityState := opacityAdjuster(fileContent)
	fontSizeAdjuster, setFontSizeState := fontSizeAdjuster(fileContent)
	fontShuffler, setFontState := fontShuffler(fileContent)

	var currentWidget ui.Drawable = themesList // default widget
	currentWidget.(*widgets.List).BorderStyle.Fg = ui.ColorYellow

	ui.Render(
		themesList,
		opacityGauge,
		fontSizeAdjuster,
		fontShuffler,
	)

	rowOne := []ui.Drawable{
		themesList,
		fontSizeAdjuster,
	}
	rowTwo := []ui.Drawable{
		opacityGauge,
	}
	rowThree := []ui.Drawable{
		fontShuffler,
	}
	widgetGrid := [][]ui.Drawable{
		rowOne,
		rowTwo,
		rowThree,
	}

	var activeRowIndex, activeColumnIndex int
	var activeWidget ui.Drawable

	e := setThemeState() // default active widget

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

			switch aw.Title {
			case "Themes":
				e = setThemeState()
			case "Font Sizes":
				e = setFontSizeState()
			case "Fonts":
				e = setFontState()
			}
		case *widgets.Gauge:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			switch aw.Title {
			case "Opacity":
				e = SetOpacityState()
			}
		}
	}
}
