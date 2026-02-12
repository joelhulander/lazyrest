package internal

import (
	"os"
	"path/filepath"

	"github.com/rivo/tview"

	"codeberg.org/joelhulander/lazyrest/internal/ui"
)

type LazyRestApp struct {
	tviewApp *tview.Application
}

func NewApp() *LazyRestApp {
	var rootDir string
	if dir, exists := os.LookupEnv("XDG_DATA_HOME"); exists {
		rootDir = filepath.Join(dir + "/lazyrest")
	} 

	tree := ui.NewFileTree(rootDir)

	app := &LazyRestApp{
		tview.NewApplication().SetRoot(tree.GetView(), true).EnableMouse(true).SetTitle("lazyrest"),
	}

	return app
}

func (a *LazyRestApp) Run() error {
	return a.tviewApp.Run()
}
