package ui

import (
	ui "github.com/gizak/termui/v3"
)

type UIWidget interface {
	InitWidget(*string)
	SetState() string
	GetWidget() ui.Drawable
}
