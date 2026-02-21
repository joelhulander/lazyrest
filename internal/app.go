package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/joelhulander/lazyrest/internal/ui"
)

type LazyRestApp struct {
	tviewApp *tview.Application
	explorer *ui.CollectionsExplorer
	requestUrlBar *ui.RequestUrlBar
	layout *ui.Layout
}

func NewApp(rootDir string) *LazyRestApp {
	tviewApp := tview.NewApplication()

	ui.SetupStyle()

	var layout *ui.Layout

	focusWorkspaceGrid := func() {
		tviewApp.SetFocus(layout.WorkspaceGrid)
	}

	var app *LazyRestApp
	tree := ui.NewCollectionsExplorer(rootDir, errorLogger)
	urlBar := ui.NewRequestUrlBar(focusWorkspaceGrid)
	textView := ui.NewResponseViewer()
	dropDown := ui.NewMethodDropDown(focusWorkspaceGrid)
	layout = ui.NewLayout(tviewApp, tree, dropDown, urlBar, textView, messagesLogger)

	app = &LazyRestApp{
		tviewApp: tviewApp,
		explorer: tree,
		requestUrlBar: urlBar,
		layout: layout,
	}

	return app
}

func (a *LazyRestApp) Run() error {
	a.tviewApp.
		EnableMouse(true).
		SetTitle("lazyrest").
		SetRoot(a.layout.GetView(), true).
		SetFocus(a.explorer.GetView())
	a.SetKeybindings()
	return a.tviewApp.Run()
}

func (a *LazyRestApp) SetKeybindings() {
	a.tviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			a.tviewApp.SetFocus(a.layout.FocusNext())
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				a.tviewApp.SetFocus(a.explorer.GetView())
			case '2':
				a.tviewApp.SetFocus(a.layout.WorkspaceGrid)
			}
		}
		return event
	})
}
