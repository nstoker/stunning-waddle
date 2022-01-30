package tutorials

import (
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func ManagingWindows() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("Hello, World of Fyne!")

	updateTime(clock)
	w.SetContent(clock)
	w.SetMaster()

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.Show()

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third window"))
		w3.Resize(fyne.NewSize(200, 50))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	a.Run()
}

// Previously declared
// func updateTime(clock *widget.Label) {
// 	formatted := time.Now().Format(time.RFC1123)
// 	clock.SetText(formatted)
// }
