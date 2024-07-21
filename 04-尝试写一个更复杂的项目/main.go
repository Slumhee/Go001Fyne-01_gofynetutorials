package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct{
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("Gofer go go go!")
	output, entry, btn, count, incrementbtn := myApp.makeUI()

	w.SetContent(container.NewVBox(output, entry, btn, count, incrementbtn))
	w.Resize(fyne.Size{Width: 1500, Height: 900})
	w.ShowAndRun()
}

func (app *App) makeUI() (*widget.Label, *widget.Entry, *widget.Button, *widget.Label, *widget.Button){
	output := widget.NewLabel("Hello World!")
	entry := widget.NewEntry()
	btn := widget.NewButton("Click!", func() {
		app.output.SetText(entry.Text)
	})

	app.output = output

	btn.Importance = widget.HighImportance
	number := 0
	count := widget.NewLabel(fmt.Sprintf("Current Number : %d", number))
	incrementbtn := widget.NewButton("Click Again!", func() {
		number++
		count.SetText(fmt.Sprintf("Current Number : %d", number))
	})
	incrementbtn.Importance = widget.DangerImportance
	return output, entry, btn, count, incrementbtn
}