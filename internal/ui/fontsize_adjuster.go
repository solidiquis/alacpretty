package ui

import (
	"fmt"
	"strconv"

	"github.com/solidiquis/alacpretty/internal/yamlconf"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type FontsizeAdjuster struct {
	// public
	Widget *widgets.List

	// private
	x1, y1, x2, y2 int
	yamlConfig     *string
}

func NewFontsizeAdjuster(x1, y1, x2, y2 int) *FontsizeAdjuster {
	return &FontsizeAdjuster{
		Widget: widgets.NewList(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}

func (fa *FontsizeAdjuster) GetWidget() ui.Drawable {
	var widget ui.Drawable = fa.Widget
	return widget
}

func (fa *FontsizeAdjuster) InitWidget(fileContent *string) {
	fa.yamlConfig = fileContent

	rows := make([]string, 27)
	for i := 6; i < 33; i++ {
		rows[i-6] = fmt.Sprintf("%d", i)
	}

	fa.Widget.Title = "Size"
	fa.Widget.Rows = rows
	fa.Widget.TextStyle = ui.NewStyle(ui.ColorYellow)
	fa.Widget.WrapText = false
	fa.Widget.SetRect(fa.x1, fa.y1, fa.x2, fa.y2)
	fa.Widget.BorderStyle.Fg = ui.ColorWhite

	currentFS := yamlconf.CurrentFontSize(fileContent)
	for i, fontSize := range fa.Widget.Rows {
		fs, _ := strconv.Atoi(fontSize)

		if currentFS == fs {
			fa.Widget.SelectedRow = i
		}
	}
}

func (fa *FontsizeAdjuster) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents
		switch e.ID {
		case "<C-c>", "q", "H", "J", "K", "L":
			return e.ID
		case "j", "<Down>":
			fa.Widget.ScrollDown()
		case "k", "<Up>":
			fa.Widget.ScrollUp()
		}

		newFS, _ := strconv.Atoi(fa.Widget.Rows[fa.Widget.SelectedRow])

		yamlconf.ChangeFontSize(fa.yamlConfig, newFS)
		yamlconf.ApplyChanges(*(fa.yamlConfig))

		ui.Render(fa.Widget)
	}
}
