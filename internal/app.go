package internal

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/joelhulander/lazyrest/internal/ui"
	"github.com/joelhulander/lazyrest/internal/appctx"
)

type App struct {
	ctx *appctx.Context
	explorer *ui.CollectionsExplorer
	workspaceGrid *ui.WorkspaceGrid
	layout *ui.Layout
}

func NewApp(rootDir string) *App {
	var layout *ui.Layout
	var explorer *ui.CollectionsExplorer
	var workspaceGrid *ui.WorkspaceGrid

	application := tview.NewApplication()
	
	focusExplorer := func() {
		application.SetFocus(explorer.GetView())
	}

	focusWorkspaceGrid := func() {
		application.SetFocus(workspaceGrid.GetView())
	}

	focusRequestPanelPage := func() {
		pageName, page := workspaceGrid.GetRequestPanel().GetPages().GetFrontPage()
		switch pageName {
		case "Params", "Headers":
			table := page.(*tview.Table)
			table.SetSelectable(true, true)
			application.SetFocus(table)
		}
	}

	focusResponsePanel := func() {
		application.SetFocus(workspaceGrid.GetResponsePanel().GetView())
	}

	focusRequestPanel := func() {
		application.SetFocus(workspaceGrid.GetRequestPanel().GetView())
	}

	onFileSelected := func (path string) {
		// fileContent, err := os.ReadFile(path)

		// if err != nil {
		// 	logger.Error("error occurred", "err", err)
		// }

		// responsePanel.GetView().SetText(string(fileContent), false)
	}

	ctx := &appctx.Context {
		App:                   application,
		FocusWorkspace:        focusWorkspaceGrid,
		FocusRequestPanelPage: focusRequestPanelPage,
		FocusResponsePanel:    focusResponsePanel,
		FocusRequestPanel:     focusRequestPanel,
		FocusExplorer:         focusExplorer,
		OnFileSelected:        onFileSelected,
	}

	ui.SetupStyle()

	explorer = ui.NewCollectionsExplorer(ctx, rootDir)
	workspaceGrid = ui.NewWorkspaceGrid(ctx)
	layout = ui.NewLayout(explorer, workspaceGrid, logger)

	app := &App{
		ctx: ctx,
		explorer: explorer,
		workspaceGrid: workspaceGrid,
		layout: layout,
	}

	return app
}

func (a *App) Run() error {
	a.ctx.App.
		EnableMouse(true).
		SetTitle("lazyrest").
		SetRoot(a.layout.GetView(), true).
		SetFocus(a.explorer.GetView())
	a.SetKeybindings()
	return a.ctx.App.Run()
}

func (a *App) SetKeybindings() {
	a.ctx.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		currentFocus := a.ctx.App.GetFocus()

		switch currentFocus.(type) {
		case *tview.InputField:
			return event
		}

		// switch currentFocus {
		// case a.workspaceGrid.GetRequestPanel().GetView(), 
		// 	a.workspaceGrid.GetResponsePanel().GetView(), 
		// 	a.workspaceGrid.GetUrlBar().GetView(), (*tview.InputField):
		// }


		if a.workspaceGrid.GetRequestPanel().HasFocus() || a.workspaceGrid.GetResponsePanel().HasFocus() {
			return event
		}

		switch event.Key() {
		case tcell.KeyTab:
			a.ctx.App.SetFocus(a.FocusNext())
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				logger.Info("Setting focus to explorer")
				a.ctx.App.SetFocus(a.explorer.GetView())
			case '2':
				logger.Info("Setting focus to workspace")
				a.ctx.App.SetFocus(a.workspaceGrid.GetView())
			case '3':
				logger.Info("Setting focus to request panel")
				a.ctx.App.SetFocus(a.workspaceGrid.GetRequestPanel().GetView())
			case '4':
				a.ctx.App.SetFocus(a.workspaceGrid.GetResponsePanel().GetView())
				logger.Info("Setting focus to response panel")
			}
		}
		return event
	})
}

func (a *App) FocusNext() tview.Primitive {
	currentFocus := a.ctx.App.GetFocus()

	if currentFocus == a.explorer.GetView() {
		return a.workspaceGrid.GetView()
	}

	return a.explorer.GetView()
}

