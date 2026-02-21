package ui

import "github.com/rivo/tview"

var responseViewer *ResponseViewer

type ResponseViewer struct {
	view *tview.TextView
}

func NewResponseViewer() *ResponseViewer {
	textView := tview.NewTextView()

	responseViewer = &ResponseViewer{
		view: textView,
	}

	textView.SetBorder(true)

	return responseViewer
}

func (r *ResponseViewer) GetView() tview.Primitive {
	return r.view
}

