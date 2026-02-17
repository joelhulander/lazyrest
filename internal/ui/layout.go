package ui

import "github.com/rivo/tview"

type Layout struct {
	root *tview.Flex
}

func NewLayout(tree *FileTree, input *UrlField) *Layout{
	flex := tview.NewFlex().
		AddItem(tree.GetView(), 0, 1, false).
		AddItem(input.GetView(), 0, 5, false)

	layout := &Layout {
		root: flex,
	}

	return layout
}

func (layout *Layout) GetView() tview.Primitive {
	return layout.root
}
