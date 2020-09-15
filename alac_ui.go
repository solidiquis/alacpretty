package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func themeShuffler(fileContent *string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = []string{
		"Argonaut",
		"Ayu Dark",
		"Ayu Mirage",
		"After Glow",
		"Base16 Default Dark",
		"Blood Moon",
		"Default Theme",
		"Solarized Light",
	}
	themesList.TextStyle = ui.NewStyle(ui.ColorYellow)
	themesList.WrapText = false
	themesList.SetRect(0, 0, 25, 8)

	ui.Render(themesList)
	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>", "<Enter>":
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

// func determineCurrentTheme(fileContent *string) {
// TODO
// }
