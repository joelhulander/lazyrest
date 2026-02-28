package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

type RequestPanel struct {
	ctx *appctx.Context
	view *tview.Flex
	pages *tview.Pages
	buttons []*tview.Button
	paramsButton *tview.Button
	headersButton *tview.Button
	paramsPage *tview.Table
	headersPage *tview.TextArea
}

func NewRequestPanel(ctx *appctx.Context) *RequestPanel {
	parent := tview.NewFlex().SetDirection(0)
	parent.SetBorder(true)
	parent.SetTitle(" [3] Request ").SetTitleAlign(0).SetBorderPadding(1, 0, 1, 1)
	pages := tview.NewPages()

	panel := &RequestPanel {
		ctx: ctx,
		view: parent,
		pages: pages,
	}

	paramsButton, headersButton := panel.NewButtons()

	buttonsFlex := tview.NewFlex().AddItem(paramsButton, 0, 1, false).AddItem(headersButton, 0, 1, false)
	buttonsFlex.SetDirection(1)

	textArea := tview.NewTextArea()
	textArea.SetText("Params", true)

	paramsTable := panel.NewTable("params")


	textArea2 := tview.NewTextArea()

	panel.buttons = append(panel.buttons, paramsButton, headersButton)
	panel.paramsPage = paramsTable
	panel.headersPage = textArea2

	panel.setActiveButton(paramsButton)
	pages.SetBorderPadding(1,1,1,1)
	pages.AddPage("Params", paramsTable, true, true)
	pages.AddPage("Headers", textArea2, true, false)
	pages.SwitchToPage("Params")

	panel.addPageSettings(textArea)
	panel.addPageSettings(textArea2)

	parent.AddItem(buttonsFlex, 1, 1, false).AddItem(pages, 0, 1, false)

	parent.SetFocusFunc(focusColorFunc(parent.Box))

	parent.SetInputCapture(panel.inputCapture)

	return panel
}

func (p *RequestPanel) NewTable(page string) *tview.Table{
	table := tview.NewTable()
	table.SetFixed(1, 1)

	table.SetCell(0, 0, &tview.TableCell{
		Text: "Name",
		Color: tcell.ColorRed,
		NotSelectable: true,
		Expansion: 1,
	})
	table.SetCell(0, 1, &tview.TableCell{
		Text: "Value",
		Color: tcell.ColorRed,
		NotSelectable: true,
		Expansion: 1,
	})
	table.SetSelectedFunc(func(row int, column int) {
		name := tview.NewInputField()
		name.SetLabel("New value: ")
		name.SetDoneFunc(func (key tcell.Key) {
			switch key {
			case tcell.KeyEnter:
				name.GetText()
				table.GetCell(row, column).SetText(name.GetText())
			}
			p.view.RemoveItem(name)
			p.ctx.FocusRequestPanelPage()
		})
		p.view.AddItem(name, 1, 1, false)
		p.ctx.App.SetFocus(name)
	})

	table.SetCell(1, 0, tview.NewTableCell("Name").SetSelectedStyle(tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack)))
	table.SetCell(1, 1, tview.NewTableCell("Value").SetSelectedStyle(tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack)))

	return table
}

func (p *RequestPanel) setActiveButton(button *tview.Button) {
	for _, b := range p.buttons {
		b.SetStyle(tcell.Style{})
	}
	button.SetStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple))
	button.SetActivatedStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorPurple))
}

func (p *RequestPanel) addPageSettings(page tview.Primitive) {
}

func (p *RequestPanel) inputCapture(event *tcell.EventKey) *tcell.EventKey {
	log.Info("In request panel input capture")
	currentFocus := p.ctx.App.GetFocus()

	pageName, frontPage := p.pages.GetFrontPage()

	isTable := false
	switch pageName {
	case "Params": 
		isTable = true
	}

	switch currentFocus {
	case frontPage:
		log.Info("Focus is on page")
		switch event.Key() {
		case tcell.KeyEscape:
			if isTable {
				frontPage.(*tview.Table).SetSelectable(false, false)
			}
			p.ctx.App.SetFocus(p.view)
			return nil
		case tcell.KeyTab:
			return nil
		}
	case p.view:
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				p.ctx.FocusExplorer()
			case '2':
				p.ctx.FocusWorkspace()
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
			blurColorFunc(p.view.Box)()
		case tcell.KeyEscape:
			p.ctx.FocusWorkspace()
			blurColorFunc(p.view.Box)()
		}
	}

	if currentFocus == p.paramsButton || currentFocus == p.headersButton {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'i':
				p.ctx.FocusRequestPanelPage()
			case 'p':
				p.setActiveButton(p.paramsButton)
				p.pages.SwitchToPage("Params")
				return nil
			case 'h':
				p.setActiveButton(p.headersButton)
				p.pages.SwitchToPage("Headers")
				return nil
			}
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

func (p *RequestPanel) GetPages() *tview.Pages {
	return p.pages
}

func (p *RequestPanel) GetView() *tview.Flex {
	return p.view
}
