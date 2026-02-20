package ui

import "github.com/rivo/tview"

var responseView *ResponseView

type ResponseView struct {
	textArea *tview.TextView
}

func NewResponseView() *ResponseView {
	textView := tview.NewTextView()

	responseView = &ResponseView{
		textArea: textView,
	}

	return responseView
}

func (r *ResponseView) GetView() tview.Primitive {
	return r.textArea
}

