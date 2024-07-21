package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.NewWithID("01")
	customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
	a.Settings().SetTheme(&myTheme{font: customFont})

	w := a.NewWindow("高端检测器")

	btn := widget.NewButton("检测是否开机", func() {
		progress := widget.NewProgressBarInfinite()
		progressContainer := container.NewVBox(progress)

		loadingDialog := dialog.NewCustom("正在检测...", "取消", progressContainer, w)
		loadingDialog.Show()

	go func(){	time.Sleep(20 * time.Second)
		loadingDialog.Hide()
		dialog.ShowInformation("结果", "电脑是开机的", w)}()
	})

	w.SetContent(container.NewVBox(btn))
	w.Resize(fyne.Size{Width: 114, Height: 514})
	w.CenterOnScreen()
	w.ShowAndRun()
}