package ui

import (
	"log"
	"slices"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Layout struct {
	app *tview.Application
	root            *tview.Grid
	WorkspaceGrid     *tview.Grid
	activeBox *tview.Box
	messagesLogger *log.Logger
	requestUrlBar *RequestUrlBar
	methods *tview.DropDown
	focusOrder []*tview.Box
}

var layout *Layout

func NewLayout(app *tview.Application, explorer *CollectionsExplorer, dropDown *MethodDropDown, urlBar *RequestUrlBar, responseViewer *ResponseViewer, messagesLogger *log.Logger) *Layout {

	flex := tview.NewFlex().AddItem(dropDown.view, 6, 1, false).AddItem(urlBar.field, 0, 1, false)

	area := tview.NewTextArea()
	area.SetBorder(true)
	rFlex := tview.NewFlex().AddItem(area, 1, 1, false)
	rFlex.SetBorder(true).SetTitle(" Request ")

	rrFlex := tview.NewFlex().AddItem(responseViewer.view, 1, 1, false)
	rrFlex.SetBorder(true).SetTitle(" Response ")

	workspaceGrid := tview.NewGrid().
		SetRows(2, 0).
		SetColumns(0, 0).
		AddItem(flex, 0, 0, 1, 2, 0, 0, false).
		// Horizontal split
		AddItem(rFlex, 1, 0, 1, 2, 0, 0, false).
		AddItem(rrFlex, 2, 0, 1, 2, 0, 0, false)
		// Vertical split
		// AddItem(rFlex, 1, 0, 2, 1, 0, 0, false).
		// AddItem(rrFlex, 1, 1, 2, 1, 0, 0, false)

	screenGrid := tview.NewGrid().
		SetColumns(30, 0).
		AddItem(explorer.view, 0, 0, 1, 1, 0, 0, false).
		AddItem(workspaceGrid, 0, 1, 1, 1, 0, 0, false)

	workspaceGrid.
		SetBorder(true).
		SetTitle("[2] Workspace ").
		SetTitleAlign(0).
		SetBorderPadding(1,0,1,1)

	focusOrder := []*tview.Box{explorer.view.Box, workspaceGrid.Box}

	layout := &Layout{
		app: app,
		root: screenGrid,
		WorkspaceGrid: workspaceGrid,
		activeBox: explorer.view.Box,
		messagesLogger: messagesLogger,
		requestUrlBar: urlBar,
		methods: dropDown.view,
		focusOrder: focusOrder,
	}

	workspaceGrid.SetFocusFunc(layout.focusColorFunc(workspaceGrid.Box))
	workspaceGrid.SetBlurFunc(layout.blurColorFunc(workspaceGrid.Box))
	explorer.view.SetFocusFunc(layout.focusColorFunc(explorer.view.Box))
	explorer.view.SetBlurFunc(layout.blurColorFunc(explorer.view.Box))

	workspaceGrid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if app.GetFocus() == dropDown.view || dropDown.view.IsOpen() {
			return event
		}

		if layout.activeBox == workspaceGrid.Box {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case 'i':
					app.SetFocus(layout.requestUrlBar.field)
					return nil
				case 'm':
					app.SetFocus(layout.methods)
					return nil
				}
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
	})

	return layout
}

func (l *Layout) GetView() tview.Primitive {
	return l.root
}

func (l *Layout) focusColorFunc(box *tview.Box) func (){
	return func () {
		l.activeBox = box
		box.SetBorderColor(tcell.ColorGreen)
		box.SetTitleColor(tcell.ColorGreen)
	}
}

func (l *Layout) blurColorFunc(box *tview.Box) func (){
	return func () {
		box.SetBorderColor(tcell.ColorGray)
		box.SetTitleColor(tcell.ColorWhite)
	}
}

func (l *Layout) FocusNext() tview.Primitive {
	currentActiveBoxIndex := slices.Index(l.focusOrder, l.activeBox)
	nextIndex := currentActiveBoxIndex + 1

	if nextIndex >= len(l.focusOrder) {
		nextIndex = 0
	} 

	return l.focusOrder[nextIndex]
}

