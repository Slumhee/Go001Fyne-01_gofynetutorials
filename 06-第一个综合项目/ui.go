package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
   )

func (cfg *config) makeUI() (*widget.Entry, *widget.RichText){
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")

	cfg.Edit = edit
	cfg.Preview = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (cfg *config) createMenu(win fyne.Window){
	open := fyne.NewMenuItem("打开...", cfg.openFunc(win))
	save := fyne.NewMenuItem("保存", cfg.saveFunc(win))
	cfg.MenuItem = save
	cfg.MenuItem.Disabled = true

	saveAs := fyne.NewMenuItem("另存为...", cfg.saveAsFunc(win))

	fileMenu := fyne.NewMenu("文件", open, save, saveAs)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}
