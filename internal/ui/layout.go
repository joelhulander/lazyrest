package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Layout struct {
	screenGrid    *tview.Grid
	workspaceGrid *tview.Grid
}

var layout *Layout

func NewLayout(explorer *CollectionsExplorer, workspaceGrid *WorkspaceGrid) *Layout {
	screenGrid := tview.NewGrid().
		SetColumns(30, 0).
		AddItem(explorer.view, 0, 0, 1, 1, 0, 0, false).
		AddItem(workspaceGrid.view, 0, 1, 1, 1, 0, 0, false)

	layout = &Layout{
		screenGrid: screenGrid,
	}

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

