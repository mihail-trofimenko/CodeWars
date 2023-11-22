package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("test")

	label := widget.NewLabel("Test label")

	w.SetContent(container.NewVBox(
		label,
	))
	w.ShowAndRun()
}
