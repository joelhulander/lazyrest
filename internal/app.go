package internal

import (
	"os"
	"path/filepath"

	"github.com/rivo/tview"

	"codeberg.org/joelhulander/lazyrest/internal/ui"
)

type LazyRestApp struct {
	tviewApp *tview.Application
	fileTree *ui.FileTree
	urlField *ui.UrlField
	layout *ui.Layout
}

func NewApp() *LazyRestApp {
	ui.SetupStyle()

	var rootDir string
	if dir, exists := os.LookupEnv("XDG_DATA_HOME"); exists {
		rootDir = filepath.Join(dir + "/lazyrest")
	} 

	tree := ui.NewFileTree(rootDir)
	input := ui.NewUrlField()
	layout := ui.NewLayout(tree, input)
	
	tviewApp := tview.
		NewApplication()

	app := &LazyRestApp{
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
		SetFocus(a.urlField.GetView())
	return a.tviewApp.Run()
}

