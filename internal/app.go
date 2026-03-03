package internal

import (
	"net/url"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/joelhulander/lazyrest/internal/client"
	"github.com/joelhulander/lazyrest/internal/ui"
)

type App struct {
	ctx           *appctx.Context
	explorer      *ui.CollectionsExplorer
	workspaceGrid *ui.WorkspaceGrid
	layout        *ui.Layout
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

	onFileSelected := func(path string) {
		// fileContent, err := os.ReadFile(path)

		// if err != nil {
		// 	logger.Error("error occurred", "err", err)
		// }

		// responsePanel.GetView().SetText(string(fileContent), false)
	}

	syncUrlParams := func() {
		baseUrl := workspaceGrid.GetUrlBar().GetText()
		if baseUrl == "" {
			return
		}

		baseUrl = stripQueryParams(baseUrl)
		params := buildQueryString(workspaceGrid.GetRequestPanel().GetParams())

		workspaceGrid.GetUrlBar().SetText(baseUrl + params)
	}

	client := client.NewClient()

	ctx := &appctx.Context{
		App:                   application,
		Client:                client,
		Logger:                logger,
		FocusWorkspace:        focusWorkspaceGrid,
		FocusRequestPanelPage: focusRequestPanelPage,
		FocusResponsePanel:    focusResponsePanel,
		FocusRequestPanel:     focusRequestPanel,
		FocusExplorer:         focusExplorer,
		OnFileSelected:        onFileSelected,
		SyncUrlParams:         syncUrlParams,
	}

	ui.SetupStyle()

	explorer = ui.NewCollectionsExplorer(ctx, rootDir)
	workspaceGrid = ui.NewWorkspaceGrid(ctx)
	layout = ui.NewLayout(explorer, workspaceGrid)

	app := &App{
		ctx:           ctx,
		explorer:      explorer,
		workspaceGrid: workspaceGrid,
		layout:        layout,
	}

	return app
}

func (a *App) Run() error {
	a.ctx.Logger.Info("starting lazyrest")
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

		if a.workspaceGrid.GetRequestPanel().HasFocus() || a.workspaceGrid.GetResponsePanel().HasFocus() {
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
				a.ctx.App.SetFocus(a.workspaceGrid.GetView())
			case '3':
				a.ctx.App.SetFocus(a.workspaceGrid.GetRequestPanel().GetView())
			case '4':
				a.ctx.App.SetFocus(a.workspaceGrid.GetResponsePanel().GetView())
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

func stripQueryParams(baseUrl string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		logger.Error("error while parsing url", "err", err)
	}
	u.RawQuery = ""

	return u.String()
}

func buildQueryString(params map[string]string) string {
	sb := strings.Builder{}

	first := true
	for k, v := range params {
		if first {
			sb.WriteString("?")
		} else {
			sb.WriteString("&")
		}
		sb.WriteString(url.QueryEscape(k))
		sb.WriteString("=")
		sb.WriteString(url.QueryEscape(v))
		first = false
	}

	return sb.String()
}
