package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/joelhulander/lazyrest/internal/client"
	"github.com/rivo/tview"
)

type WorkspaceGrid struct {
	ctx           *appctx.Context
	view          *tview.Grid
	requestPanel  *RequestPanel
	responsePanel *ResponsePanel
	requestUrlBar *RequestUrlBar
	methods       *MethodDropDown
}

func NewWorkspaceGrid(ctx *appctx.Context) *WorkspaceGrid {
	urlBar := NewRequestUrlBar(ctx)
	requestPanel := NewRequestPanel(ctx)
	responsePanel := NewResponsePanel(ctx)
	dropDown := NewMethodDropDown(ctx)

	flex := tview.NewFlex().AddItem(dropDown.view, 6, 1, false).AddItem(urlBar.field, 0, 1, false)

	rFlex := tview.
		NewFlex().
		SetDirection(0).
		AddItem(requestPanel.view, 0, 1, false).
		AddItem(responsePanel.view, 0, 1, false)

	view := tview.NewGrid().
		SetRows(2, 0).
		SetColumns(0, 0).
		AddItem(flex, 0, 0, 1, 2, 0, 0, false).
		AddItem(rFlex, 1, 0, 1, 2, 0, 0, false)

	view.
		SetBorder(true).
		SetTitle(" [2] Workspace ").
		SetTitleAlign(0).
		SetBorderPadding(1, 0, 1, 1)

	grid := &WorkspaceGrid{
		ctx:           ctx,
		view:          view,
		methods:       dropDown,
		requestUrlBar: urlBar,
		requestPanel:  requestPanel,
		responsePanel: responsePanel,
	}

	view.SetFocusFunc(focusColorFunc(view.Box))
	view.SetBlurFunc(blurColorFunc(view.Box))
	view.SetInputCapture(grid.inputCapture)

	return grid
}

func (g *WorkspaceGrid) GetView() *tview.Grid {
	return g.view
}

func (g *WorkspaceGrid) GetRequestPanel() *RequestPanel {
	return g.requestPanel
}

func (g *WorkspaceGrid) GetResponsePanel() *ResponsePanel {
	return g.responsePanel
}

func (g *WorkspaceGrid) GetUrlBar() *RequestUrlBar{
	return g.requestUrlBar
}

func (g *WorkspaceGrid) GetMethodsDropdown() *MethodDropDown {
	return g.methods
}

func (g *WorkspaceGrid) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	currentFocus := g.ctx.App.GetFocus()

	switch currentFocus.(type) {
	case *tview.InputField:
		return event
	}

	if g.requestPanel.HasFocus() || g.responsePanel.HasFocus() || g.methods.HasFocusOrIsOpen() || g.requestUrlBar.HasFocus() {
		return event
	}

	switch event.Key() {
	case tcell.KeyRune:
		switch event.Rune() {
		case 'i':
			g.ctx.App.SetFocus(g.requestUrlBar.field)
			return nil
		case 'm':
			g.ctx.App.SetFocus(g.methods.view)
			return nil
		}
	case tcell.KeyEnter:
		url := g.requestUrlBar.GetText()
		_, method := g.methods.view.GetCurrentOption()

		request := client.Request {
			URL: url, 
			Method: method,
		}

		resp, err := g.ctx.Client.SendRequest(request)
		
		if err != nil {
			g.ctx.Logger.Error("Error while sending the request", "err", err)
			return nil
		}

		g.responsePanel.textView.SetText(resp.Body)


		return nil
	case tcell.KeyUp:
		return nil
	case tcell.KeyDown:
		return nil
	case tcell.KeyRight:
		return nil
	case tcell.KeyLeft:
		return nil
	}

	return event

}
