package tutorials

import (
	"testing"

	"fyne.io/fyne/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	if out.Text != "Hello, World!" {
		t.Error("incorrect initial variable")
	}

	test.Type(in, "Tulip")
	if out.Text != "Hello, Tulip!" {
		t.Error("incorrect user greeting")
	}
}
