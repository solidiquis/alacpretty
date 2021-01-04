package ui

import (
	ui "github.com/gizak/termui/v3"
)

type UIWidget interface {
	GetWidget() ui.Drawable
	InitWidget(*string)
	SetState() string
	ToggleActive()
}
