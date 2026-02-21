package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SetupStyle() {
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

