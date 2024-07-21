package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	Edit *widget.Entry
	Preview *widget.RichText
	CurrentFile fyne.URI
	MenuItem *fyne.MenuItem
	BaseTitle string
}

var cfg config

var filter = storage.NewExtensionFileFilter([]string{".md", ".MD"})

func main() {
	a := app.NewWithID("01")
	customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
	a.Settings().SetTheme(&myTheme{font: customFont})
	// fmt.Println(cfg.CurrentFile)
	w := a.NewWindow("Markdown编辑器")
	cfg.BaseTitle = "Markdown编辑器"


	edit, preview := cfg.makeUI()
	cfg.createMenu(w)
	w.SetContent(container.NewHSplit(edit, preview))
	w.Resize(fyne.Size{Width: 800, Height: 600})
	w.CenterOnScreen()
	w.ShowAndRun()
}


