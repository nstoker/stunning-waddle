package tutorials

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Clock() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("Hello, World of Fyne!")

	updateTime(clock)
	w.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format(time.RFC1123)
	clock.SetText(formatted)
}
