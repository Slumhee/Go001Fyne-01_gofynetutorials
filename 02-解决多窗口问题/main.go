package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w1 := a.NewWindow("Window 1")
	w1.SetContent(container.NewVBox(widget.NewLabel("Test01")))

	w2 := a.NewWindow("Window 2")
	w2.SetContent(container.NewVBox(widget.NewLabel("Test02")))

	w1.Show()
	w2.Show()
	a.Run()
}