package ui

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var layout *Layout

type Layout struct {
	app *tview.Application
	root            *tview.Flex
	CollectionsFlex *tview.Flex
	RequestFlex     *tview.Flex
	activeBox *tview.Box
	currentFocus    int
	messagesLogger *log.Logger
	urlField *UrlField
}

func NewLayout(app *tview.Application, tree *FileTree, input *UrlField, textView *ResponseView, messagesLogger *log.Logger) *Layout {
	collectionsFlex := tview.NewFlex()
	requestFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	flex := collectionsFlex.
		AddItem(tree.GetView(), 0, 1, false).
		AddItem(requestFlex.
			AddItem(input.GetView(), 0, 1, false).
			AddItem(textView.GetView(), 0, 17, false),
			0, 2, false)

	requestFlex.SetBorder(true).SetTitle("Request")

	layout := &Layout{
		app: app,
		root:            flex,
		CollectionsFlex: collectionsFlex,
		RequestFlex:     requestFlex,
		currentFocus:    0,
		messagesLogger: messagesLogger,
		urlField: input,
	}

	for i := range layout.root.GetItemCount() {
		item := layout.root.GetItem(i)
		switch v := item.(type) {
		case *tview.Flex:
			v.SetFocusFunc(layout.focusColorFunc(v.Box, v.GetTitle()))
			v.SetBlurFunc(layout.blurColorFunc(v.Box))
		case *tview.TreeView:
			v.SetFocusFunc(layout.focusColorFunc(v.Box, v.GetTitle()))
			v.SetBlurFunc(layout.blurColorFunc(v.Box))
		}
	}

	requestFlex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if layout.activeBox == requestFlex.Box {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case 'i':
					// requestFlex.SetFocus(a.layout.FocusItem(a.fileTree.GetView()))
					app.SetFocus(layout.urlField.input)
					return nil
				}
			}
		}
		return event
	})

	return layout
}

func (l *Layout) GetView() tview.Primitive {
	return l.root
}

func (l *Layout) focusColorFunc(box *tview.Box, title string) func (){
	return func () {
		l.messagesLogger.Println("Moving focus to: " + title)
		l.activeBox = box
		box.SetBorderColor(tcell.ColorGreen)
		// l.messagesLogger.Printf("Focusing on box: %v\n", box)
	}
}

func (l *Layout) blurColorFunc(box *tview.Box) func (){
	return func () {
		// l.messagesLogger.Printf("Leaving focus on box: %v\n", box)
		box.SetBorderColor(tcell.ColorGray)
	}
}

func (l *Layout) FocusNext() tview.Primitive {
	if l.currentFocus == l.root.GetItemCount()-1 {
		l.currentFocus = 0
	} else {
		l.currentFocus += 1
	}

	return l.root.GetItem(l.currentFocus)
}

func (l *Layout) FocusItem(t tview.Primitive) tview.Primitive {
	for i := range l.root.GetItemCount() {
		if t == l.root.GetItem(i) {
			return l.root.GetItem(i)
		}
	}
	return nil
}

