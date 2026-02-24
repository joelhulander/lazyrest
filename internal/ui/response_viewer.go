package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ResponseView struct {
	view *tview.TextView
}

var responseView *ResponseView

func NewResponseView(onEscape func ()) *ResponseView {
	view := tview.NewTextView()

	responseView := &ResponseView{
		view: view,
	}

	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			onEscape()
		}
		return event
	})

	view.SetTitle(" Response ").SetBorder(true)

	view.SetFocusFunc(focusColorFunc(view.Box))
	view.SetBlurFunc(blurColorFunc(view.Box))

	return responseView
}

func (r *ResponseView) GetView() *tview.TextView {
	return r.view
}

