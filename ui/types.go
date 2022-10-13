package ui

import (
	"fyne.io/fyne/v2"
	"pixl/apptype"
	"pixl/pxcanvas"
	"pixl/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
