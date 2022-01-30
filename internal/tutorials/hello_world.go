package tutorials

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func HelloWorld() {

	fmt.Println("Starting Hello World app.")
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello, World of Fyne!"))

	w.ShowAndRun()
}
