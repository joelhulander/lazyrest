package appctx

import (
	"log/slog"

	"github.com/joelhulander/lazyrest/internal/client"
	"github.com/rivo/tview"
)

type Context struct {
	App                   *tview.Application
	Client                *client.Client
	Logger                *slog.Logger
	FocusWorkspace        func()
	FocusRequestPanel     func()
	FocusResponsePanel    func()
	FocusRequestPanelPage func()
	FocusExplorer         func()
	OnFileSelected        func(path string)
	SyncUrlParams         func()
}
