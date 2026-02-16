package internal

import (
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"codeberg.org/joelhulander/lazyrest/internal/ui"
)

type LazyRestApp struct {
	tviewApp *tview.Application
	fileTree *tview.TreeView
	urlField *tview.InputField
}

func NewApp() *LazyRestApp {
	setupStyle()

	var rootDir string
	if dir, exists := os.LookupEnv("XDG_DATA_HOME"); exists {
		rootDir = filepath.Join(dir + "/lazyrest")
	} 

	tree := ui.NewFileTree(rootDir)
	tree.GetView().SetBorder(true)

	input := tview.NewInputField()
	
	flexBox := tview.NewFlex().
		AddItem(tree.GetView(), 0, 1, false).
		AddItem(input, 0, 5, false)

	app := &LazyRestApp{
		tviewApp: tview.NewApplication().SetRoot(flexBox, true).EnableMouse(true).SetTitle("lazyrest"),
		fileTree: tree.GetView(),
		urlField: input,
	}

	return app
}

func (a *LazyRestApp) Run() error {
	a.urlField.SetBorder(true)
	a.tviewApp.SetFocus(a.urlField)
	return a.tviewApp.Run()
}

func setupStyle() {
	tview.Borders.TopLeft = '╭'
	tview.Borders.TopRight = '╮'
	tview.Borders.BottomLeft = '╰'
	tview.Borders.BottomRight = '╯'

	tview.Borders.TopLeftFocus = '╭'
	tview.Borders.TopRightFocus = '╮'
	tview.Borders.BottomLeftFocus = '╰'
	tview.Borders.BottomRightFocus = '╯'

	tview.Borders.Horizontal = '─'
	tview.Borders.Vertical = '│'

	tview.Borders.HorizontalFocus = '─'
	tview.Borders.VerticalFocus = '│'

	tview.Styles.PrimitiveBackgroundColor = tcell.NewRGBColor(25, 23, 36)
	tview.Styles.ContrastBackgroundColor = tcell.NewRGBColor(33, 32, 46)
	tview.Styles.MoreContrastBackgroundColor = tcell.NewRGBColor(42, 39, 63)
	tview.Styles.BorderColor = tcell.NewRGBColor(110, 106, 134)
	tview.Styles.TitleColor = tcell.NewRGBColor(224, 222, 244)
	tview.Styles.GraphicsColor = tcell.NewRGBColor(156, 207, 216)
	tview.Styles.PrimaryTextColor = tcell.NewRGBColor(224, 222, 244)
	tview.Styles.SecondaryTextColor = tcell.NewRGBColor(246, 193, 119)
	tview.Styles.TertiaryTextColor = tcell.NewRGBColor(156, 207, 216)
	tview.Styles.InverseTextColor = tcell.NewRGBColor(25, 23, 36)
	tview.Styles.ContrastSecondaryTextColor = tcell.NewRGBColor(235, 111, 146)

}
