package main

import (
  "fmt"

  "fyne.io/fyne/v2"

  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/widget"
)

func main() {
  a:= app.New()
  customFont := fyne.NewStaticResource("NotoSansHans.ttf", loadFont("NotoSansHans-Regular.ttf"))
  a.Settings().SetTheme(&myTheme{font: customFont})
  w:= a.NewWindow("让我们一起来学习Go语言吧!")

  w.SetContent(widget.NewLabel("让我们一起来学习Go语言吧!"))
  w.ShowAndRun()

  fmt.Println("close!")
}