package ui

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/solidiquis/alacpretty/internal/themes"
	"github.com/solidiquis/alacpretty/internal/utils"
	"github.com/solidiquis/alacpretty/internal/yamlconf"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
}

func themeShuffler(fileContent *string) (*widgets.List, func() string) {
	var rows []string
	for theme := range themes.AllThemes {
		rows = append(rows, theme)
	}

	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = rows
	themesList.TextStyle = ui.NewStyle(ui.ColorYellow)
	themesList.WrapText = false
	themesList.SetRect(0, 0, 25, 8)
	themesList.BorderStyle.Fg = ui.ColorWhite

	currentTheme := yamlconf.CurrentTheme(fileContent)
	for index, theme := range themesList.Rows {
		if currentTheme == theme {
			themesList.SelectedRow = index
		}
	}

	setState := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "q", "H", "J", "K", "L":
				return e.ID
			case "j", "<Down>":
				themesList.ScrollDown()
			case "k", "<Up>":
				themesList.ScrollUp()
			}

			newTheme := themesList.Rows[themesList.SelectedRow]
			yamlconf.ChangeTheme(fileContent, newTheme)
			yamlconf.ApplyChanges(*fileContent)

			ui.Render(themesList)
		}
	}

	return themesList, setState
}

func opacityAdjuster(fileContent *string) (*widgets.Gauge, func() string) {
	opacityGauge := widgets.NewGauge()
	opacityGauge.Title = "Opacity"
	opacityGauge.SetRect(0, 8, 50, 11)
	opacityGauge.Percent = int(yamlconf.CurrentOpacity(fileContent) * 100)
	opacityGauge.BarColor = ui.ColorYellow
	opacityGauge.LabelStyle = ui.NewStyle(ui.ColorBlue)
	opacityGauge.BorderStyle.Fg = ui.ColorWhite

	setState := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			var newOpacity float64
			var tmp int

			switch e.ID {
			case "<C-c>", "q", "H", "J", "K", "L":
				return e.ID
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
				yamlconf.ChangeOpacity(fileContent, newOpacity)
				yamlconf.ApplyChanges(*fileContent)
				ui.Render(opacityGauge)
			}
		}
	}

	return opacityGauge, setState
}

func fontSizeAdjuster(fileContent *string) (*widgets.List, func() string) {
	rows := make([]string, 27)

	for i := 6; i < 33; i++ {
		rows[i-6] = fmt.Sprintf("%d", i)
	}

	fsList := widgets.NewList()
	fsList.Title = "Font Sizes"
	fsList.Rows = rows
	fsList.TextStyle = ui.NewStyle(ui.ColorYellow)
	fsList.WrapText = false
	fsList.SetRect(26, 0, 50, 8)
	fsList.BorderStyle.Fg = ui.ColorWhite

	currentFS := yamlconf.CurrentFontSize(fileContent)
	for index, fontSize := range fsList.Rows {
		fs, err := strconv.Atoi(fontSize)
		utils.Must(err)

		if currentFS == fs {
			fsList.SelectedRow = index
		}
	}

	setState := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "q", "H", "J", "K", "L":
				return e.ID
			case "j":
				fsList.ScrollDown()
			case "k":
				fsList.ScrollUp()
			}

			newFS, err := strconv.Atoi(fsList.Rows[fsList.SelectedRow])
			utils.Must(err)

			yamlconf.ChangeFontSize(fileContent, newFS)
			yamlconf.ApplyChanges(*fileContent)

			ui.Render(fsList)
		}
	}

	return fsList, setState
}

func WidgetsController(fileContent *string) {
	defer ui.Close()

	themesList, setThemeState := themeShuffler(fileContent)
	opacityGauge, SetOpacityState := opacityAdjuster(fileContent)
	fontSizeAdjuster, setFontSizeState := fontSizeAdjuster(fileContent)

	var currentWidget ui.Drawable = themesList // default widget
	currentWidget.(*widgets.List).BorderStyle.Fg = ui.ColorYellow

	ui.Render(
		themesList,
		opacityGauge,
		fontSizeAdjuster,
	)

	rowOne := []ui.Drawable{
		themesList,
		fontSizeAdjuster,
	}
	rowTwo := []ui.Drawable{
		opacityGauge,
	}
	widgetGrid := [][]ui.Drawable{
		rowOne,
		rowTwo,
	}

	var activeRowIndex, activeColumnIndex int
	var activeWidget ui.Drawable

	e := setThemeState()

	for {
		switch e {
		case "J":
			activeRowIndex--
		case "K":
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
		} else if activeColumnIndex > len(activeRow)-1 {
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
			}
		case *widgets.Gauge:
			aw.BorderStyle.Fg = ui.ColorYellow
			ui.Render(aw)
			currentWidget = aw
			e = SetOpacityState()
		}
	}
}
