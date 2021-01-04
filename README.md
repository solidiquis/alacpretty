# Alacpretty ‎️‍🌈

<a href="https://github.com/alacritty/alacritty">Alacritty</a> is a GPU-enhanced terminal emulator; Alacpretty is a Go program that leverages <a href="https://github.com/gizak/termui">termui</a> to provide Alacritty users the ability to edit their Alacritty configurations via a terminal user-interface.

<p align="center">
  <img src="https://github.com/solidiquis/alacpretty/blob/master/assets/alacpretty.gif">
</p> 

## How it works

Alacpretty components are ready to plug into your Alacritty configurations. Assuming your `alacritty.yml` file is in the appropriate location as specified in the <a href="https://github.com/alacritty/alacritty">Alacritty documentation</a>, all you'll need to do is grab the components you need from `internal/ui`, and drop them into `WidgetsController` like it is done in `cmd/alacpretty/main.go`.

```go
package main

import (
	aui "github.com/solidiquis/alacpretty/internal/ui"
	"github.com/solidiquis/alacpretty/internal/yamlconf"
)

func main() {
  // Reads the contents of alacritty.yml to a string
  content := yamlconf.ReadFileToString()
  
  // Set the position of each widget, and arrange in rows
  row1 := []aui.UIWidget{
    aui.NewThemeShuffler(0, 0, 25, 10),
    aui.NewFontsizeAdjuster(26, 0, 51, 10),
  }
  row2 := []aui.UIWidget{
    aui.NewOpacityGauge(0, 11, 51, 14),
  }

  // Controller handles rendering and navigating between widgets.
  aui.WidgetsController(
    &content,
    row1,
    row2,
  )
}
```

Here is how you navigate the UI:
- `h`, `j`, `k`, `l` and `←`, `↓`, `↑` , `→` are equivalent and are used to navigate inside the widget or change its appearance.
- `H`, `J`, `K`, `L` are used to navigate across widgets, as only one can be focused on at at time.

They way in which the focus shifts between widgets is dependent upon the way you arrange your widgets in rows.

## How to use

There are a couple ways in which you can use Alacpretty: The simplest way would be to have Go installed, clone this repo, and run `go run cmd/alacpretty/main.go` - or you can clone this repo, compile the code into binary, and stick it somewhere in your path.

## Stable widgets

1. <a href="https://github.com/solidiquis/alacpretty/blob/master/internal/ui/theme_shuffler.go">ThemeShuffler</a>
2. <a href="https://github.com/solidiquis/alacpretty/blob/master/internal/ui/opacity_gauge.go">OpacityGauge</a>
3. <a href="https://github.com/solidiquis/alacpretty/blob/master/internal/ui/fontsize_adjuster.go">FontsizeAdjuster</a>

## Work-in-progress widgets

1. FontShuffler
2. FontSearchbar
3. ThemeSearchbar
4. HelpBox
5. YamlDisplay

## Licence

<a href="https://raw.githubusercontent.com/solidiquis/alacpretty/master/LICENCE">MIT</a>

## Note from author

Thank you to everyone who has taken an interest in this little pet project of mine — I honestly didn't expect it to get so many stars! And with that said, I also apologize for how slowly I am moving as work has been keeping me extraordinarily busy. If you have any questions, concerns, or would like to contribute, don't hesitate to write up an issue or submit a PR!
