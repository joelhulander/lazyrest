package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/joelhulander/lazyrest/internal/appctx"
	"github.com/rivo/tview"
)

type RequestUrlBar struct {
	ctx *appctx.Context
	field *tview.InputField
}

var requestUrlBar *RequestUrlBar

func NewRequestUrlBar(ctx *appctx.Context) *RequestUrlBar {
	field := tview.NewInputField()
	field.
		SetPlaceholder("Enter URL").
		SetPlaceholderTextColor(tcell.ColorGray).SetLabel(" ")


	urlBar := &RequestUrlBar {
		ctx: ctx,
		field: field,
	}


	urlBar.field.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			ctx.Logger.Debug("url bar dismissed")
			ctx.FocusWorkspace()
		case tcell.KeyEnter:
			ctx.Logger.Info("url submitted", "url", urlBar.field.GetText())
			ctx.FocusWorkspace()
		}
		return event
	})

	requestUrlBar = urlBar

	return urlBar
}

func (bar *RequestUrlBar) SetText(text string) {
	bar.field.SetText(text)
}

func (bar *RequestUrlBar) GetText() string {
	return bar.field.GetText()
}

func (bar *RequestUrlBar) HasFocus() bool {
	f := bar.ctx.App.GetFocus()
	return f == bar.field
}

func (bar *RequestUrlBar) GetView() tview.Primitive {
	return bar.field
}

