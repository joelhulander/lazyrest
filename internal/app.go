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
	requestPanel *ui.RequestPanel
	responseView *ui.ResponseView
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
	tree := ui.NewCollectionsExplorer(rootDir, logger)
	urlBar := ui.NewRequestUrlBar(focusWorkspaceGrid)
	requestPanel := ui.NewRequestPanel(focusWorkspaceGrid)
	responseView := ui.NewResponseView(focusWorkspaceGrid)
	dropDown := ui.NewMethodDropDown(focusWorkspaceGrid)
	layout = ui.NewLayout(tviewApp, tree, dropDown, urlBar, requestPanel, responseView, logger)

	app = &LazyRestApp{
		tviewApp: tviewApp,
		explorer: tree,
		requestUrlBar: urlBar,
		requestPanel: requestPanel,
		responseView: responseView,
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
		currentFocus := a.tviewApp.GetFocus()
		if currentFocus == a.requestPanel.GetView() || currentFocus == a.responseView.GetView() {
			return event
		}

		switch event.Key() {
		case tcell.KeyTab:
			a.tviewApp.SetFocus(a.FocusNext())
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

func (a *LazyRestApp) FocusNext() tview.Primitive {
	logger.Debug("In FocusNext()")
	currentFocus := a.tviewApp.GetFocus()

	if currentFocus == a.explorer.GetView().Box {
		logger.Debug("Setting focus on workspace")
		return a.layout.WorkspaceGrid.Box
	}

	logger.Debug("Setting focus on explorer")
	return a.explorer.GetView().Box

}

