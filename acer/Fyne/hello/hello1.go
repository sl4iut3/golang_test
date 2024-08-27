package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	cpt := 0
	a := app.New()
	w := a.NewWindow("Hello")
	clock := widget.NewLabel("")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		clock,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
			cpt = 0
		}),
	))
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
			cpt++
			hello.SetText(fmt.Sprintf("welcome %d", cpt))
		}
	}()

	w.ShowAndRun()
}
