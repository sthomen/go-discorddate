package gui

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/container"

	"git.shangtai.net/staffan/go-discorddate/internal/dateformat"
)

// App UI context
type _context struct {
	date time.Time
	format dateformat.DateFormat
	formatWidget *widget.Select
	previewWidget *widget.Label
}

var context = _context{}

func MainWindow(date time.Time, format dateformat.DateFormat) {
	context.date = date
	context.format = format

	app := app.New()
	w := app.NewWindow("Discord Date Tool")

	context.formatWidget = widget.NewSelect(dateformat.FormatsOptions(), updateUi)
	context.previewWidget = widget.NewLabel("")

	button := widget.NewButton("Copy to clipboard", func () {
		w.Clipboard().SetContent(context.format.Render(context.date))
	})

	w.SetContent(container.NewVBox(
		context.formatWidget,
		button,
		context.previewWidget,
	))

	// Set initial value, but this has to be done after SetContent or fyne will
	// panic.
	context.formatWidget.SetSelected(context.format.Name)

	w.ShowAndRun()
}

func updateUi(_ string) {
	var index = context.formatWidget.SelectedIndex()

	format, err := dateformat.ByIndex(index)

	if err == nil {
		context.format = format
	}

	context.previewWidget.SetText(context.format.Render(context.date))
}
