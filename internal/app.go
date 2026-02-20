package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/joelhulander/lazyrest/internal/ui"
)

type LazyRestApp struct {
	tviewApp *tview.Application
	fileTree *ui.FileTree
	urlField *ui.UrlField
	layout *ui.Layout
}

func NewApp(rootDir string) *LazyRestApp {
	tviewApp := tview.NewApplication()

	ui.SetupStyle()

	var layout *ui.Layout
	var app *LazyRestApp
	tree := ui.NewFileTree(rootDir, errorLogger)
	input := ui.NewUrlField(func() {tviewApp.SetFocus(layout.RequestFlex)})
	textView := ui.NewResponseView()
	layout = ui.NewLayout(tviewApp, tree, input, textView, messagesLogger)

	app = &LazyRestApp{
		tviewApp: tviewApp,
		fileTree: tree,
		urlField: input,
		layout: layout,
	}

	return app
}

func (a *LazyRestApp) Run() error {
	a.tviewApp.
		EnableMouse(true).
		SetTitle("lazyrest").
		SetRoot(a.layout.GetView(), true).
		SetFocus(a.fileTree.GetView())
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
				a.tviewApp.SetFocus(a.layout.CollectionsFlex)
			case '2':
				a.tviewApp.SetFocus(a.layout.RequestFlex)
			}
		}
		return event
	})
}
