package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

type ResponsePanel struct {
	view *tview.Flex
}

func NewResponsePanel(ctx *appctx.Context) *ResponsePanel {
	view := tview.NewFlex()

	panel := &ResponsePanel{
		view: view,
	}

	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			ctx.FocusRequestPanel()
		case tcell.KeyEscape:
			ctx.FocusWorkspace()
		}
		return event
	})

	view.SetTitle(" [4] Response ").SetTitleAlign(0).SetBorderPadding(1, 0, 1, 1).SetBorder(true)

	view.SetFocusFunc(focusColorFunc(view.Box))
	view.SetBlurFunc(blurColorFunc(view.Box))

	return panel
}

func (r *ResponsePanel) GetView() *tview.Flex {
	return r.view
}

