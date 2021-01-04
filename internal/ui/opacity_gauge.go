package ui

import (
	"github.com/solidiquis/alacpretty/internal/yamlconf"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type OpacityGauge struct {
	// public
	Widget *widgets.Gauge

	// private
	x1, y1, x2, y2 int
	yamlConfig     *string
}

func NewOpacityGauge(x1, y1, x2, y2 int) *OpacityGauge {
	return &OpacityGauge{
		Widget: widgets.NewGauge(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}

func (og *OpacityGauge) GetWidget() ui.Drawable {
	var widget ui.Drawable = og.Widget
	return widget
}

func (og *OpacityGauge) InitWidget(fileContent *string) {
	og.yamlConfig = fileContent

	og.Widget.Title = "Opacity"
	og.Widget.SetRect(og.x1, og.y1, og.x2, og.y2)
	og.Widget.Percent = int(yamlconf.CurrentOpacity(fileContent) * 100)
	og.Widget.BarColor = ui.ColorYellow
	og.Widget.LabelStyle = ui.NewStyle(ui.ColorBlue)
	og.Widget.BorderStyle.Fg = ui.ColorWhite
}

func (og *OpacityGauge) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents

		var newOpacity float64
		var tmp int
		switch e.ID {
		case "<C-c>", "q", "H", "J", "K", "L":
			return e.ID
		case "l", "<Left>":
			tmp = og.Widget.Percent + 10
		case "h", "<Right>":
			tmp = og.Widget.Percent - 10
		default:
			tmp = 1925 // Cthulhu Fhtagn
		}

		if tmp >= 0 && tmp <= 100 {
			newOpacity = float64(tmp) / 100
			og.Widget.Percent = tmp
			yamlconf.ChangeOpacity(og.yamlConfig, newOpacity)
			yamlconf.ApplyChanges(*(og.yamlConfig))
			ui.Render(og.Widget)
		}
	}
}
