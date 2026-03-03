package appctx

import (
	"log/slog"

	"github.com/rivo/tview"
)

type Context struct {
	App                   *tview.Application
	Logger                *slog.Logger
	FocusWorkspace        func()
	FocusRequestPanel     func()
	FocusResponsePanel    func()
	FocusRequestPanelPage func()
	FocusExplorer         func()
	OnFileSelected        func(path string)
}

