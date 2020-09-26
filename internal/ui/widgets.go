package ui

import (
	"fmt"
	"strconv"

	"github.com/solidiquis/alacpretty/internal/themes"
	"github.com/solidiquis/alacpretty/internal/utils"
	"github.com/solidiquis/alacpretty/internal/yamlconf"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

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