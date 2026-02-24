package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RequestUrlBar struct {
	field *tview.InputField
}

var requestUrlBar *RequestUrlBar

func NewRequestUrlBar(onEscape func()) *RequestUrlBar {
	field := tview.NewInputField()
	field.
		SetPlaceholder("Enter URL").
		SetPlaceholderTextColor(tcell.ColorGray).SetLabel(" ")


	urlBar := &RequestUrlBar {
		field: field,
	}


	urlBar.field.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			onEscape()
		}
		return event
	})

	requestUrlBar = urlBar

	return urlBar
}

func (bar *RequestUrlBar) SetText(text string) {
	bar.field.SetText(text)
}

func (bar *RequestUrlBar) GetText() string {
	return bar.field.GetText()
}

func (bar *RequestUrlBar) GetView() tview.Primitive {
	return bar.field
}

