package tutorials

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func HelloPerson() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))

	w.ShowAndRun()
}

func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello, World!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello, " + content + "!")
	}

	return out, in
}
