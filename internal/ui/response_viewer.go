package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

type ResponsePanel struct {
	ctx *appctx.Context
	view *tview.Flex
}

func NewResponsePanel(ctx *appctx.Context) *ResponsePanel {
	view := tview.NewFlex()

	panel := &ResponsePanel{
		ctx:  ctx,
		view: view,
	}

	view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			ctx.FocusRequestPanel()
		case tcell.KeyEscape:
			ctx.FocusWorkspace()
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				ctx.FocusExplorer()
				return nil
			case '2':
				ctx.FocusWorkspace()
				return nil
			case '3':
				ctx.FocusRequestPanel()
				return nil
			}
		}
		return event
	})

	view.SetTitle(" [4] Response ").SetTitleAlign(0).SetBorderPadding(1, 0, 1, 1).SetBorder(true)

	view.SetFocusFunc(focusColorFunc(view.Box))
	view.SetBlurFunc(blurColorFunc(view.Box))

	return panel
}

func (p *ResponsePanel) HasFocus() bool {
	f := p.ctx.App.GetFocus()
	return f == p.view 
}


func (r *ResponsePanel) GetView() *tview.Flex {
	return r.view
}

