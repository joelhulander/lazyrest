package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type UrlField struct {
	input *tview.InputField
	onEscape func()
}

var urlField *UrlField

func NewUrlField(onEscape func()) *UrlField {
	input := tview.NewInputField()

	field := &UrlField {
		input: input,
		onEscape: onEscape,
	}

	field.input.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			field.input.Blur()
			onEscape()
		}
		return event
	})

	urlField = field

	return field
}

func (f *UrlField) SetText(text string) {
	f.input.SetText(text)
}

func (f *UrlField) GetText() string {
	return f.input.GetText()
}

func (f *UrlField) GetView() tview.Primitive {
	return f.input
}

