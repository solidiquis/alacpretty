package ui

import (
	"sort"

	"github.com/solidiquis/alacpretty/internal/themes"
	"github.com/solidiquis/alacpretty/internal/yamlconf"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type ThemeShuffler struct {
	// Public
	Widget *widgets.List

	// private
	x1, y1, x2, y2 int
	yamlConfig     *string
}

func NewThemeShuffler(x1, y1, x2, y2 int) *ThemeShuffler {
	return &ThemeShuffler{
		Widget: widgets.NewList(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}

func (ts *ThemeShuffler) GetWidget() ui.Drawable {
	var widget ui.Drawable = ts.Widget
	return widget
}

func (ts *ThemeShuffler) InitWidget(fileContent *string) {
	ts.yamlConfig = fileContent

	currentTheme := yamlconf.CurrentTheme(fileContent)
	rows := make([]string, len(themes.ThemeNames))

	var selectedRow int
	for i, theme := range themes.ThemeNames {
		rows[i] = theme
		if currentTheme == theme {
			selectedRow = i
		}
	}
	sort.Strings(rows)

	ts.Widget.Title = "Themes"
	ts.Widget.Rows = rows
	ts.Widget.SelectedRow = selectedRow
	ts.Widget.TextStyle = ui.NewStyle(ui.ColorYellow)
	ts.Widget.WrapText = false
	ts.Widget.SetRect(ts.x1, ts.y1, ts.x2, ts.y2)
	ts.Widget.BorderStyle.Fg = ui.ColorWhite
}

func (ts *ThemeShuffler) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents
		switch e.ID {
		case "<C-c>", "q", "H", "J", "K", "L":
			return e.ID
		case "j", "down":
			ts.Widget.ScrollDown()
		case "k", "up":
			ts.Widget.ScrollUp()
		}

		newTheme := ts.Widget.Rows[ts.Widget.SelectedRow]
		yamlconf.ChangeTheme(ts.yamlConfig, newTheme)
		yamlconf.ApplyChanges(*(ts.yamlConfig))

		ui.Render(ts.GetWidget())
	}
}
