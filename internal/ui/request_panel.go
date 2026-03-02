package ui

import (
	"fmt"

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
	paramsPage    *tview.Table
	headersPage   *tview.Table
}

func NewRequestPanel(ctx *appctx.Context) *RequestPanel {
	parent := tview.NewFlex().SetDirection(0)
	parent.SetBorder(true)
	parent.SetTitle(" [3] Request ").SetTitleAlign(0).SetBorderPadding(1, 0, 1, 1)
	pages := tview.NewPages()

	panel := &RequestPanel{
		ctx:   ctx,
		view:  parent,
		pages: pages,
	}

	paramsButton, headersButton := panel.NewButtons()

	buttonsFlex := tview.NewFlex().AddItem(paramsButton, 0, 1, false).AddItem(headersButton, 0, 1, false)
	buttonsFlex.SetDirection(1)

	paramsTable := panel.NewTable()
	headersTable := panel.NewTable()

	panel.buttons = append(panel.buttons, paramsButton, headersButton)
	panel.paramsPage = paramsTable
	panel.headersPage = headersTable

	panel.setActiveButton(paramsButton)
	pages.SetBorderPadding(1, 1, 1, 1)
	pages.AddPage("Params", paramsTable, true, true)
	pages.AddPage("Headers", headersTable, true, false)
	pages.SwitchToPage("Params")

	parent.AddItem(buttonsFlex, 1, 1, false).AddItem(pages, 0, 1, false)
	parent.SetFocusFunc(focusColorFunc(parent.Box))
	parent.SetBlurFunc(blurColorFunc(parent.Box))
	parent.SetInputCapture(panel.inputCapture)

	return panel
}

func (p *RequestPanel) NewTable() *tview.Table {
	table := tview.NewTable()
	table.SetFixed(1, 1)

	table.SetCell(0, 0, &tview.TableCell{Text: "Name", Color: tcell.ColorRed, NotSelectable: true, Expansion: 1})
	table.SetCell(0, 1, &tview.TableCell{Text: "Value", Color: tcell.ColorRed, NotSelectable: true, Expansion: 1})
	table.SetCell(1, 0, tview.NewTableCell("Name").
		SetStyle(tcell.Style{}.Foreground(tcell.ColorGray)).
		SetSelectedStyle(tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack)))
	table.SetCell(1, 1, tview.NewTableCell("Value").
		SetStyle(tcell.Style{}.Foreground(tcell.ColorGray)).
		SetSelectedStyle(tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack)))

	table.SetSelectedFunc(func(row int, column int) {
		input := tview.NewInputField()
		input.SetLabel("New value: ")
		input.SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				table.GetCell(row, column).SetText(input.GetText()).SetStyle(tcell.Style{}.Foreground(tcell.ColorWhite))
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

func (p *RequestPanel) setActiveButton(button *tview.Button) {
	for _, b := range p.buttons {
		b.SetStyle(tcell.Style{})
	}
	button.SetStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple))
	button.SetActivatedStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple))
}

func (p *RequestPanel) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	currentFocus := p.ctx.App.GetFocus()
	log.Info("Current focus is on request panel", "type", fmt.Sprintf("%T", currentFocus))

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
		}

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
				p.ctx.FocusRequestPanelPage()
				return nil
			case 'p':
				p.setActiveButton(p.paramsButton)
				p.pages.SwitchToPage("Params")
				return nil
			case 'h':
				p.setActiveButton(p.headersButton)
				p.pages.SwitchToPage("Headers")
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

func (p *RequestPanel) NewButtons() (params *tview.Button, headers *tview.Button) {
	params = tview.NewButton("Params")
	headers = tview.NewButton("Headers")
	p.paramsButton = params
	p.headersButton = headers
	return
}

func (p *RequestPanel) HasFocus() bool {
	f := p.ctx.App.GetFocus()
	return f == p.view || f == p.pages || f == p.paramsPage || f == p.headersPage
}

func (p *RequestPanel) GetPages() *tview.Pages {
	return p.pages
}

func (p *RequestPanel) GetView() *tview.Flex {
	return p.view
}
