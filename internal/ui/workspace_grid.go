package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

type WorkspaceGrid struct {
	ctx *appctx.Context
	view          *tview.Grid
	requestPanel  *RequestPanel
	responsePanel  *ResponsePanel
	requestUrlBar *RequestUrlBar
	methods       *MethodDropDown
}

func NewWorkspaceGrid(ctx *appctx.Context,
	dropDown *MethodDropDown, urlBar *RequestUrlBar,
	requestPanel *RequestPanel, responsePanel *ResponsePanel) *WorkspaceGrid {

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
		ctx: ctx,
		view:          view,
		methods:       dropDown,
		requestUrlBar: requestUrlBar,
		requestPanel:  requestPanel,
		responsePanel:  responsePanel,
	}

	view.SetInputCapture(grid.inputCapture)

	return grid
}

func (g *WorkspaceGrid) GetView() *tview.Grid {
	return g.view
}

func (g *WorkspaceGrid) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	log.Info("In workspace input capture")
	currentFocus := g.ctx.App.GetFocus()

	if currentFocus == g.methods.view || g.methods.view.IsOpen() {
		return event
	}

	if currentFocus == g.view {
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
			g.ctx.App.SetFocus(g.requestPanel.view)
		case tcell.KeyUp:
			return nil
		case tcell.KeyDown:
			return nil
		case tcell.KeyRight:
			return nil
		case tcell.KeyLeft:
			return nil
		}
	}

	return event

}

