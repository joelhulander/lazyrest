package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MethodDropDown struct {
	view *tview.DropDown
}

func NewMethodDropDown(onEscape func()) *MethodDropDown {
	view := tview.NewDropDown().AddOption(" GET ", nil).AddOption(" POST ", nil).AddOption(" PUT ", nil).SetCurrentOption(0)
	view.SetFieldBackgroundColor(tcell.ColorPurple).SetFieldTextColor(tcell.ColorBlack)

	selectedStyle := tcell.Style{}
	unselectedStyle := tcell.Style{}
	view.SetListStyles(unselectedStyle.Background(tcell.ColorGray).Foreground(tcell.ColorBlack), selectedStyle.Background(tcell.ColorRed).Foreground(tcell.ColorBlack))

	dropDown := &MethodDropDown {
		view: view,
	}

	dropDown.view.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			onEscape()
		case tcell.KeyEnter:
			return event
		case tcell.KeyUp:
			return event
		case tcell.KeyDown:
			return event
		case tcell.KeyLeft:
			return event
		case tcell.KeyRight:
			return event
		default:
			return nil
		}
		return event
	})


	return dropDown
}
