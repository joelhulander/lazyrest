package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type RequestPanel struct {
	view *tview.Flex
	pages *tview.Pages
	paramsButton *tview.Button
	headersButton *tview.Button
}

var requestPanel *RequestPanel

func NewRequestPanel(onEscape func ()) *RequestPanel {
	parent := tview.NewFlex().SetDirection(0)
	parent.SetBorder(true)
	parent.SetTitle(" Request ")
	pages := tview.NewPages()

	panel := &RequestPanel {
		view: parent,
		pages: pages,

	}

	panel.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			onEscape()
		}
		return event
	})

	paramsButton, headersButton := panel.NewButtons()

	buttonsFlex := tview.NewFlex().AddItem(paramsButton, 0, 1, false).AddItem(headersButton, 0, 1, false)
	buttonsFlex.SetDirection(1)

	textView := tview.NewTextView()
	textView.SetBackgroundColor(tcell.ColorRed)

	pages.SetBorderPadding(1,1,1,1)
	pages.AddPage("params", textView, true, true)
	pages.AddPage("headers", tview.NewTextArea().SetBackgroundColor(tcell.ColorBlue), true, false)

	parent.AddItem(buttonsFlex, 1, 1, false).AddItem(pages, 0, 1, false)

	parent.SetFocusFunc(focusColorFunc(parent.Box))
	parent.SetBlurFunc(blurColorFunc(parent.Box))

	return panel
}


func (p *RequestPanel) NewButtons() (params *tview.Button, headers *tview.Button) {
	// buttonsStyle := tcell.Style {}
	// buttonsStyle.Underline(tcell.UnderlineStyleSolid)
	// headersButton.SetStyle(buttonsStyle)
	params = tview.NewButton("Params")
	params.SetFocusFunc(func () {
		p.pages.SwitchToPage("params")
	})
	headers = tview.NewButton("Headers")
	headers.SetFocusFunc(func () {
		p.pages.SwitchToPage("headers")
	})

	p.paramsButton = params
	p.headersButton = headers
	// paramsButton.SetStyle(buttonsStyle)

	return 

}

func (p *RequestPanel) GetView() *tview.Flex {
	return p.view
}
