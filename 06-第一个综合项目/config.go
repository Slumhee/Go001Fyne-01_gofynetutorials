package main

import (
	// "fmt"
	"io"
	"strings"
   
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
   )

func (cfg *config) saveAsFunc(win fyne.Window) func(){
	return func(){
		saveDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil{
				dialog.ShowError(err, win)
				return
			}

			if write == nil{
				return
			}

			if !strings.HasSuffix(strings.ToLower(write.URI().String()), ".md"){
				dialog.ShowInformation("错误", "必须是.md扩展名", win)
				return
			}

			write.Write([]uint8(cfg.Edit.Text))
			
			cfg.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(cfg.BaseTitle + "-" + write.URI().Name())

			cfg.MenuItem.Disabled = false
		}, win)
		saveDialog.SetFileName("未命名.md")
		saveDialog.SetFilter(filter)
		saveDialog.Show()
	}
}

func (cfg *config) openFunc(win fyne.Window) func() {
	return func(){
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil{
				dialog.ShowError(err, win)
				return
			}
			if read == nil{
				return
			}

			data, err := io.ReadAll(read)

			if err != nil{
				dialog.ShowError(err, win)
				return
			}

			defer read.Close()

			cfg.Edit.SetText(string(data))

			cfg.CurrentFile = read.URI()
			// fmt.Println(cfg.CurrentFile)
			win.SetTitle(cfg.BaseTitle + "-" + read.URI().Name())

			cfg.MenuItem.Disabled = false
		}, win)
		openDialog.SetFilter(filter)
		openDialog.Show()
	}
}

func (cfg *config) saveFunc(win fyne.Window) func(){
	return func(){
		if cfg.CurrentFile != nil{
			write, err := storage.Writer(cfg.CurrentFile)
			if err !=nil{
				dialog.ShowError(err, win)
				return
			}

			write.Write([]byte(cfg.Edit.Text))
			defer write.Close()
		}
	}
}