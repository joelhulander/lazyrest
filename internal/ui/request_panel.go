package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

type RequestPanel struct {
	ctx           *appctx.Context
	view          *tview.Flex
	pages         *tview.Pages
	buttons       []*tview.Button
	paramsButton  *tview.Button
	headersButton *tview.Button
	paramsTable   *tview.Table
	headersTable  *tview.Table
}

func NewRequestPanel(ctx *appctx.Context) *RequestPanel {
	parent := newParent()
	pages := tview.NewPages()

	panel := &RequestPanel{
		ctx:   ctx,
		view:  parent,
		pages: pages,
	}

	paramsButton, headersButton := panel.newButtons()

	buttonsFlex := tview.NewFlex().AddItem(paramsButton, 0, 1, false).AddItem(headersButton, 0, 1, false)
	buttonsFlex.SetDirection(1)

	paramsTable := panel.newTable()
	headersTable := panel.newTable()

	panel.buttons = append(panel.buttons, paramsButton, headersButton)
	panel.paramsTable = paramsTable
	panel.headersTable = headersTable
	panel.setActiveButton(paramsButton)

	panel.setupPages(paramsTable, headersTable)

	parent.AddItem(buttonsFlex, 1, 1, false).AddItem(pages, 0, 1, false)
	parent.SetInputCapture(panel.inputCapture)

	return panel
}

func (p *RequestPanel) GetPages() *tview.Pages {
	return p.pages
}

func (p *RequestPanel) GetView() *tview.Flex {
	return p.view
}

func (p *RequestPanel) HasFocus() bool {
	f := p.ctx.App.GetFocus()
	return f == p.view || f == p.pages || f == p.paramsTable || f == p.headersTable
}

func (p *RequestPanel) GetParams() map[string]string {
	return p.getTableData(p.paramsTable)
}

func (p *RequestPanel) GetHeaders() map[string]string {
	return p.getTableData(p.headersTable)
}

func (p *RequestPanel) getTableData(table *tview.Table) map[string]string {
	data := map[string]string{}

	for i := range table.GetRowCount() {
		if i == 0 {
			continue
		}
		name := table.GetCell(i, 0).Text
		if name == "" {
			continue
		}
		data[name] = table.GetCell(i, 1).Text
	}

	return data

}

func (p *RequestPanel) newButtons() (params *tview.Button, headers *tview.Button) {
	params = tview.NewButton("Params")
	headers = tview.NewButton("Headers")
	p.paramsButton = params
	p.headersButton = headers
	return
}

func (p *RequestPanel) newTable() *tview.Table {
	table := tview.NewTable()
	table.SetFixed(1, 1)

	table.SetCell(0, 0, &tview.TableCell{Text: "Key", Color: tcell.ColorRed, NotSelectable: true, Expansion: 1})
	table.SetCell(0, 1, &tview.TableCell{Text: "Value", Color: tcell.ColorRed, NotSelectable: true, Expansion: 1})

	table.SetSelectedFunc(func(row int, column int) {
		input := tview.NewInputField()
		if column == 0 {
			input.SetLabel("New key: ")
		} else {
			input.SetLabel("New value: ")
		}

		input.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				oldValue := table.GetCell(row, column).Text
				newValue := input.GetText()
				table.GetCell(row, column).SetText(newValue).SetStyle(tcell.Style{}.Foreground(tcell.ColorWhite))
				p.syncParamsToUrl()
				p.ctx.Logger.Info("cell edited", "row", row, "col", column, "old", oldValue, "new", newValue)
			}
			p.view.RemoveItem(input)
			p.ctx.FocusRequestPanelPage()
		})
		p.view.AddItem(input, 1, 1, false)
		p.ctx.App.SetFocus(input)
	})

	table.SetFocusFunc(focusColorFunc(p.view.Box))
	table.SetBlurFunc(blurColorFunc(p.view.Box))

	return table
}

func (p *RequestPanel) syncParamsToUrl() {
	p.ctx.SyncUrlParams()
}

func (p *RequestPanel) setActiveButton(button *tview.Button) {
	for _, b := range p.buttons {
		b.SetStyle(tcell.Style{})
	}
	button.SetStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple))
	button.SetActivatedStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple))
}

func (p *RequestPanel) setupPages(paramsTable *tview.Table, headersTable *tview.Table) {
	p.pages.SetBorderPadding(1, 1, 1, 1)
	p.pages.AddPage("Params", paramsTable, true, true)
	p.pages.AddPage("Headers", headersTable, true, false)
	p.pages.SwitchToPage("Params")
}

func (p *RequestPanel) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	currentFocus := p.ctx.App.GetFocus()

	_, frontPage := p.pages.GetFrontPage()
	table, isTable := frontPage.(*tview.Table)

	switch currentFocus {
	case frontPage:
		switch event.Key() {
		case tcell.KeyEscape:
			if isTable {
				table.SetSelectable(false, false)
			}
			p.ctx.App.SetFocus(p.view)
			return nil
		case tcell.KeyTab:

		case tcell.KeyRune:
			switch event.Rune() {
			case 'a':
				rowCount := table.GetRowCount()
				p.ctx.Logger.Info("add row", "count", rowCount)
				p.newRow(table, rowCount)
				table.Select(rowCount, 0)
				p.syncParamsToUrl()
				return nil
			case 'd':
				row, _ := table.GetSelection()
				p.ctx.Logger.Info("selected row", "row", row)
				table.RemoveRow(row)
				rowCount := table.GetRowCount()
				if row == rowCount {
					table.Select(rowCount - 1, 0)
				}
				if rowCount == 1 {
					p.ctx.FocusRequestPanel()
				}
				p.syncParamsToUrl()
				return nil
			}
		}
		return event

	case p.view:
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				p.ctx.FocusExplorer()
				return nil
			case '2':
				p.ctx.FocusWorkspace()
				return nil
			case '4':
				p.ctx.FocusResponsePanel()
				return nil
			case 'i':
				if table.GetRowCount() == 1 {
					return nil
				}
				p.ctx.FocusRequestPanelPage()
				return nil
			case 'a':
				rowCount := table.GetRowCount()
				p.ctx.Logger.Info("add row", "count", rowCount)
				p.newRow(table, rowCount)
				p.ctx.FocusRequestPanelPage()
				table.Select(rowCount, 0)
				p.syncParamsToUrl()
				return nil
			case 'p':
				p.setActiveButton(p.paramsButton)
				p.pages.SwitchToPage("Params")
				p.ctx.Logger.Debug("request tab switched", "tab", "params")
				return nil
			case 'h':
				p.setActiveButton(p.headersButton)
				p.pages.SwitchToPage("Headers")
				p.ctx.Logger.Debug("request tab switched", "tab", "headers")
				return nil
			}
		case tcell.KeyTab:
			p.ctx.FocusResponsePanel()
			return nil
		case tcell.KeyEscape:
			p.ctx.FocusWorkspace()
			return nil
		}
	}

	return event
}

func (p *RequestPanel) newRow(table *tview.Table, rowCount int) {
	table.SetCell(rowCount, 0, tview.NewTableCell("").
		SetStyle(tcell.Style{}.Foreground(tcell.ColorGray)).
		SetSelectedStyle(tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack)))
	table.SetCell(rowCount, 1, tview.NewTableCell("").
		SetStyle(tcell.Style{}.Foreground(tcell.ColorGray)).
		SetSelectedStyle(tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack)))
}

func newParent() *tview.Flex {
	parent := tview.NewFlex().SetDirection(0)
	parent.SetBorder(true)
	parent.SetTitle(" [3] Request ").SetTitleAlign(0).SetBorderPadding(1, 0, 1, 1)
	parent.SetFocusFunc(focusColorFunc(parent.Box))
	parent.SetBlurFunc(blurColorFunc(parent.Box))

	return parent
}
