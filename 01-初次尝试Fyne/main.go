package main

import(
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gofer go go go!")

	w.SetContent(widget.NewLabel("Gopher Gogogo!"))
	w.Show()
	a.Run()

	fmt.Println("closed!")
}