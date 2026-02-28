package ui

import (
	"log/slog"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Layout struct {
	screenGrid *tview.Grid
	workspaceGrid *tview.Grid
}

var log *slog.Logger
var layout *Layout

func NewLayout(explorer *CollectionsExplorer, workspaceGrid *WorkspaceGrid, logger *slog.Logger) *Layout {

	log = logger

	layout = &Layout{
		screenGrid: nil,
		workspaceGrid: workspaceGrid.view,
	}

	layout.screenGrid = tview.NewGrid().
		SetColumns(30, 0).
		AddItem(explorer.view, 0, 0, 1, 1, 0, 0, false).
		AddItem(layout.workspaceGrid, 0, 1, 1, 1, 0, 0, false)

	layout.workspaceGrid.SetFocusFunc(focusColorFunc(layout.workspaceGrid.Box))
	layout.workspaceGrid.SetBlurFunc(blurColorFunc(layout.workspaceGrid.Box))
	explorer.view.SetFocusFunc(focusColorFunc(explorer.view.Box))
	explorer.view.SetBlurFunc(blurColorFunc(explorer.view.Box))

	return layout
}

func (l *Layout) GetView() tview.Primitive {
	return l.screenGrid
}

func (l *Layout) GetWorkspaceView() *tview.Grid{
	return l.workspaceGrid
}

func focusColorFunc(box *tview.Box) func (){
	return func () {
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

