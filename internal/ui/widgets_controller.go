package ui

import (
	"log"
	"math"

	ui "github.com/gizak/termui/v3"
)

func init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
}

func WidgetsController(fileContent *string, rows ...[]UIWidget) {
	defer ui.Close()

	numWidgets := 0
	for _, row := range rows {
		numWidgets += len(row)
	}

	widgets := make([]UIWidget, numWidgets)
	drawables := make([]ui.Drawable, numWidgets) // underlying termui drawables

	i := 0
	for _, row := range rows {
		for _, widget := range row {
			widget.InitWidget(fileContent)
			drawables[i] = widget.GetWidget()
			widgets[i] = widget
			i++
		}
	}

	ui.Render(drawables...)

	// First widget -> default widget
	activeWidget := widgets[0]
	activeWidget.ToggleActive()

	activeRowIndex, activeColumnIndex := 0, 0

	e := activeWidget.SetState()
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

		// Deactivate current active widget
		activeWidget.ToggleActive()

		if activeRowIndex < 0 {
			absActiveRowIndex := int(math.Abs(float64(activeRowIndex)))
			activeRowIndex = len(rows) - absActiveRowIndex
		} else if activeRowIndex > len(rows)-1 {
			activeRowIndex = 0
		}
		activeRow := rows[activeRowIndex]

		if activeColumnIndex < 0 {
			absActiveColumnIndex := int(math.Abs(float64(activeColumnIndex)))
			activeColumnIndex = len(activeRow) - absActiveColumnIndex
		} else if activeColumnIndex >= len(activeRow) {
			activeColumnIndex = 0
		}
		activeColumn := activeRow[activeColumnIndex]

		// Set new active widget
		activeWidget = activeColumn
		activeWidget.ToggleActive()

		e = activeWidget.SetState()
	}
}
