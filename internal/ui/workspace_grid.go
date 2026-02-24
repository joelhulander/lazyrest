package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type WorkspaceGrid struct {
	app           *tview.Application
	view          *tview.Grid
	requestPanel  *RequestPanel
	responseView  *ResponseView
	requestUrlBar *RequestUrlBar
	methods       *MethodDropDown
}

func NewWorkspaceGrid(app *tview.Application,
	dropDown *MethodDropDown, urlBar *RequestUrlBar,
	requestPanel *RequestPanel, responseView *ResponseView) *WorkspaceGrid {

	flex := tview.NewFlex().AddItem(dropDown.view, 6, 1, false).AddItem(urlBar.field, 0, 1, false)

	rFlex := tview.
		NewFlex().
		SetDirection(0).
		AddItem(requestPanel.view, 0, 1, false).
		AddItem(responseView.view, 0, 1, false)

	view := tview.NewGrid().
		SetRows(2, 0).
		SetColumns(0, 0).
		AddItem(flex, 0, 0, 1, 2, 0, 0, false).
		AddItem(rFlex, 1, 0, 1, 2, 0, 0, false)

	view.
		SetBorder(true).
		SetTitle("[2] Workspace ").
		SetTitleAlign(0).
		SetBorderPadding(1, 0, 1, 1)

	grid := &WorkspaceGrid{
		app:           app,
		view:          view,
		methods:       dropDown,
		requestUrlBar: requestUrlBar,
		requestPanel:  requestPanel,
		responseView:  responseView,
	}

	view.SetInputCapture(grid.inputCapture)

	return grid
}

func (g *WorkspaceGrid) GetView() *tview.Grid {
	return g.view
}

func (g *WorkspaceGrid) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	currentFocus := g.app.GetFocus()

	if currentFocus == g.methods.view || g.methods.view.IsOpen() {
		return event
	}

	if currentFocus == g.requestPanel.view {
		g.handleRequestPanelKeys(event)
		return event
	}

	if currentFocus == g.responseView.view {
		g.handleResponseViewerKeys(event)
		return event
	}

	if currentFocus == g.view {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'i':
				g.app.SetFocus(g.requestUrlBar.field)
				return nil
			case 'm':
				g.app.SetFocus(g.methods.view)
				return nil
			}
		case tcell.KeyEnter:
			g.app.SetFocus(g.requestPanel.view)
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

func (g *WorkspaceGrid) handleRequestPanelKeys(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyTab:
		g.app.SetFocus(g.responseView.view)
	case tcell.KeyRune:
		switch event.Rune() {
		case 'p':
			g.app.SetFocus(g.requestPanel.paramsButton)
		case 'h':
			g.app.SetFocus(g.requestPanel.headersButton)
		}

	}
}

func (g *WorkspaceGrid) handleResponseViewerKeys(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyTab:
		g.app.SetFocus(g.requestPanel.view)
	}
}

