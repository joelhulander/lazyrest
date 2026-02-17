package ui

import (
	"github.com/rivo/tview"
)

type UrlField struct {
	input *tview.InputField
}


func NewUrlField() *UrlField {
	input := tview.NewInputField()
	input.SetBorder(true)

	field := &UrlField {
		input: input,
	}

	return field
}

func (field *UrlField) GetView() tview.Primitive {
	return field.input
}
