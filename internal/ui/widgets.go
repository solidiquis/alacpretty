package ui

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"github.com/solidiquis/alacpretty/internal/themes"
	"github.com/solidiquis/alacpretty/internal/utils"
	"github.com/solidiquis/alacpretty/internal/yamlconf"

	"github.com/flopp/go-findfont"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"

	"golang.org/x/sys/unix"
)

const (
	columnOneStart = 0
	columnOneEnd   = 40
)

var (
	termWidth  int
	termHeight int
)

func init() {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	utils.Must(err)

	termWidth = int(ws.Col)
	termHeight = int(ws.Row)
}

func themeSearchbox() (*widgets.Paragraph, func() string) {
	var x1, x2, y1, y2 int = columnOneStart, columnOneEnd, 0, 3

	searchbox := widgets.NewParagraph()
	searchbox.Title = "Search Theme"
	searchbox.Text = " "
	searchbox.SetRect(x1, y1, x2, y2)
	searchbox.BorderStyle.Fg = ui.ColorWhite

	processKey := func(eventID string) string {
		switch eventID {
		case "<Space>":
			return ""
		default:
			return eventID
		}
	}

	setState := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "q", "H", "J", "K", "L":
				return e.ID
			case "<Backspace>":
				if len(searchbox.Text) == 0 {
					continue
				}
				searchbox.Text = string([]byte(searchbox.Text)[:len(searchbox.Text)-1])
			default:
				searchbox.Text = searchbox.Text + processKey(e.ID)
			}
			ui.Render(searchbox)
		}
	}
	return searchbox, setState
}

func ThemeShufflerDeprecated(fileContent *string) (*widgets.List, func() string) {
	var x1, x2, y1, y2 int = columnOneStart, columnOneEnd, 4, 15

	var rows []string
	for theme := range themes.AllThemes {
		rows = append(rows, theme)
	}
	sort.Strings(rows)

	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = rows
	themesList.TextStyle = ui.NewStyle(ui.ColorYellow)
	themesList.WrapText = false
	themesList.SetRect(x1, y1, x2, y2)
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
			case "j":
				themesList.ScrollDown()
			case "k":
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
	var x1, x2, y1, y2 int = columnOneEnd + 1, termWidth, 0, 3

	opacityGauge := widgets.NewGauge()
	opacityGauge.Title = "Opacity"
	opacityGauge.SetRect(x1, y1, x2, y2)
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
	var x1, x2, y1, y2 int = 31, 40, 20, 31

	rows := make([]string, 27)

	for i := 6; i < 33; i++ {
		rows[i-6] = fmt.Sprintf("%d", i)
	}

	fsList := widgets.NewList()
	fsList.Title = "Size"
	fsList.Rows = rows
	fsList.TextStyle = ui.NewStyle(ui.ColorYellow)
	fsList.WrapText = false
	fsList.SetRect(x1, y1, x2, y2)
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

func fontSearchbox() (*widgets.Paragraph, func() string) {
	var x1, x2, y1, y2 int = columnOneStart, columnOneEnd, 16, 19

	searchbox := widgets.NewParagraph()
	searchbox.Title = "Search Font"
	searchbox.Text = " "
	searchbox.SetRect(x1, y1, x2, y2)
	searchbox.BorderStyle.Fg = ui.ColorWhite

	processKey := func(eventID string) string {
		switch eventID {
		case "<Space>":
			return ""
		default:
			return eventID
		}
	}

	setState := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "q", "H", "J", "K", "L":
				return e.ID
			case "<Backspace>":
				if len(searchbox.Text) == 0 {
					continue
				}
				searchbox.Text = string([]byte(searchbox.Text)[:len(searchbox.Text)-1])
			default:
				searchbox.Text = searchbox.Text + processKey(e.ID)
			}
			ui.Render(searchbox)
		}
	}
	return searchbox, setState
}

func fontShuffler(fileContent *string) (*widgets.List, func() string) {
	var x1, x2, y1, y2 int = columnOneStart, 30, 20, 31

	allFonts := findfont.List()
	for i, fontPath := range allFonts {
		font := strings.TrimSuffix(fontPath, filepath.Ext(fontPath))
		tmp := strings.Split(font, "/Fonts/")
		allFonts[i] = tmp[len(tmp)-1]
	}
	fontsList := widgets.NewList()
	fontsList.Title = "Fonts"
	fontsList.Rows = allFonts
	fontsList.TextStyle = ui.NewStyle(ui.ColorYellow)
	fontsList.WrapText = false
	fontsList.SetRect(x1, y1, x2, y2)
	fontsList.BorderStyle.Fg = ui.ColorWhite

	currentFont := yamlconf.CurrentFont(fileContent)

	for index, font := range fontsList.Rows {
		if currentFont == font {
			fontsList.SelectedRow = index
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
				fontsList.ScrollDown()
			case "k":
				fontsList.ScrollUp()
			}

			newFont := fontsList.Rows[fontsList.SelectedRow]
			yamlconf.ChangeFont(fileContent, newFont)
			yamlconf.ApplyChanges(*fileContent)

			ui.Render(fontsList)
		}
	}
	return fontsList, setState
}

func showYaml() (*widgets.Paragraph, chan string, func()) {
	var x1, x2, y1, y2 int = columnOneEnd + 1, termWidth, 4, termHeight

	box := widgets.NewParagraph()
	box.Title = "alacritty.yml"
	box.SetRect(x1, y1, x2, y2)
	box.BorderStyle.Fg = ui.ColorWhite

	ch := make(chan string, 1)

	// Use in goroutine
	// TODO: Make idempotent
	setState := func() {
		for {
			text := yamlconf.ReadFileToString()
			switch <-ch {
			case "Themes", "Search Theme":
				box.Text = yamlconf.TrimYamlPrefix(text, "theme")
			case "Fonts", "Search Font", "Size":
				box.Text = yamlconf.TrimYamlPrefix(text, "font")
			case "Opacity":
				box.Text = yamlconf.TrimYamlPrefix(text, "opacity")
			default:
				box.Text = ""
			}
			ui.Render(box)
		}
	}
	return box, ch, setState
}

func helpBox() *widgets.Paragraph {
	var x1, x2, y1, y2 int = columnOneStart, columnOneEnd, 32, termHeight

	box := widgets.NewParagraph()
	box.Title = "Help"
	box.Text = ""
	box.SetRect(x1, y1, x2, y2)
	box.BorderStyle.Fg = ui.ColorWhite
	return box
}
