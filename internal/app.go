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
	requestPanel *ui.RequestPanel
	responsePanel *ui.ResponsePanel
	layout *ui.Layout
}

func NewApp(rootDir string) *App {
	var layout *ui.Layout
	var explorer *ui.CollectionsExplorer
	var requestPanel *ui.RequestPanel
	var responsePanel *ui.ResponsePanel

	application := tview.NewApplication()
	
	focusExplorer := func() {
		application.SetFocus(explorer.GetView())
	}

	focusWorkspaceGrid := func() {
		application.SetFocus(layout.GetWorkspaceView())
	}

	focusRequestPanelPage := func() {
		pageName, _ := requestPanel.GetPages().GetFrontPage()
		switch pageName {
		case "Params":
			requestPanel.GetPages().GetPage(pageName).(*tview.Table).SetSelectable(true, true)
		}
		application.SetFocus(requestPanel.GetPages())
	}

	focusResponsePanel := func() {
		application.SetFocus(responsePanel.GetView())
	}

	focusRequestPanel := func() {
		application.SetFocus(requestPanel.GetView())
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
	urlBar := ui.NewRequestUrlBar(ctx)
	requestPanel = ui.NewRequestPanel(ctx)
	responsePanel = ui.NewResponsePanel(ctx)
	dropDown := ui.NewMethodDropDown(ctx)
	workspaceGrid := ui.NewWorkspaceGrid(ctx, dropDown, urlBar, requestPanel, responsePanel)
	layout = ui.NewLayout(explorer, workspaceGrid, logger)

	app := &App{
		ctx: ctx,
		explorer: explorer,
		requestPanel: requestPanel,
		responsePanel: responsePanel,
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
		if currentFocus == a.requestPanel.GetView() || currentFocus == a.responsePanel.GetView() {
			return event
		}

		switch event.Key() {
		case tcell.KeyTab:
			a.ctx.App.SetFocus(a.FocusNext())
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				a.ctx.App.SetFocus(a.explorer.GetView())
			case '2':
				a.ctx.App.SetFocus(a.layout.GetWorkspaceView())
			case '3':
				a.ctx.App.SetFocus(a.requestPanel.GetView())
			}
		}
		return event
	})
}

func (a *App) FocusNext() tview.Primitive {
	currentFocus := a.ctx.App.GetFocus()

	if currentFocus == a.explorer.GetView() {
		return a.layout.GetWorkspaceView()
	}

	return a.explorer.GetView()
}

