package ui

import (
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Layout struct {
	app *tview.Application
	screenGrid *tview.Grid
	WorkspaceGrid *tview.Grid
	focusOrder []*tview.Box
}

var log *slog.Logger
var layout *Layout

func NewLayout(app *tview.Application, explorer *CollectionsExplorer, 
	dropDown *MethodDropDown, urlBar *RequestUrlBar, 
	requestPanel *RequestPanel, responseView *ResponseView, logger *slog.Logger) *Layout {

	log = logger

	layout = &Layout{
		app: app,
		screenGrid: nil,
		WorkspaceGrid: nil,
		focusOrder: nil,
	}

	layout.WorkspaceGrid = NewWorkspaceGrid(app, dropDown, urlBar, requestPanel, responseView).view

	layout.screenGrid = tview.NewGrid().
		SetColumns(30, 0).
		AddItem(explorer.view, 0, 0, 1, 1, 0, 0, false).
		AddItem(layout.WorkspaceGrid, 0, 1, 1, 1, 0, 0, false)

	layout.focusOrder = []*tview.Box{explorer.GetView().Box, layout.WorkspaceGrid.Box}

	layout.WorkspaceGrid.SetFocusFunc(focusColorFunc(layout.WorkspaceGrid.Box))
	layout.WorkspaceGrid.SetBlurFunc(blurColorFunc(layout.WorkspaceGrid.Box))
	explorer.view.SetFocusFunc(focusColorFunc(explorer.view.Box))
	explorer.view.SetBlurFunc(blurColorFunc(explorer.view.Box))

	return layout
}

func (l *Layout) GetView() tview.Primitive {
	return l.screenGrid
}

func focusColorFunc(box *tview.Box) func (){
	return func () {
		log.Debug("New focus", "box", box.GetTitle())
		box.SetBorderColor(tcell.ColorGreen)
		box.SetTitleColor(tcell.ColorGreen)
	}
}

func blurColorFunc(box *tview.Box) func (){
	return func () {
		box.SetBorderColor(tcell.ColorGray)
		box.SetTitleColor(tcell.ColorWhite)
	}
}

