package appctx

import "github.com/rivo/tview"

type Context struct {
	App *tview.Application
	FocusWorkspace func ()
	FocusRequestPanel func ()
	FocusResponsePanel func ()
	FocusRequestPanelPage func ()
	FocusExplorer func ()
	OnFileSelected func (path string)
}

