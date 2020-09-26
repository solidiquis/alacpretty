package main

import (
	alacui "github.com/solidiquis/alacpretty/internal/ui"
	"github.com/solidiquis/alacpretty/internal/yamlconf"
)

func main() {
	content := yamlconf.ReadFileToString()
	alacui.WidgetsController(&content)
}
