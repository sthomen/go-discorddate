package gui

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"

	"git.shangtai.net/staffan/go-discorddate/internal/dateformat"
)

// App UI context
type _context struct {
	date time.Time
	format dateformat.DateFormat
	datePickerWidget *DatePicker
	formatWidget *widget.Select
	previewWidget *widget.Label
}

var context = _context{}

func MainWindow(date time.Time, format dateformat.DateFormat) {
	context.date = date
	context.format = format

	app := app.New()
	w := app.NewWindow("Discord Date Tool")

	context.datePickerWidget = NewDatePicker(func (when time.Time) {
		context.date = when
		updateUi()
	})

	context.formatWidget = widget.NewSelect(dateformat.FormatsOptions(), func(_ string) { updateUi() })
	context.previewWidget = widget.NewLabel("")

	button := widget.NewButton("Copy to clipboard", func () {
		w.Clipboard().SetContent(context.format.Render(context.date))
	})

	w.SetContent(container.NewVBox(
		context.datePickerWidget,
		context.formatWidget,
		button,
		context.previewWidget,
	))

	// Set initial values, this has to be done after SetContent or fyne will
	// panic on SetSelected().
	context.formatWidget.SetSelected(context.format.Name)
	context.datePickerWidget.SetDateTime(date)

	w.ShowAndRun()
}

func updateUi() {
	var index = context.formatWidget.SelectedIndex()

	format, err := dateformat.ByIndex(index)

	if err == nil {
		context.format = format
	}

	context.previewWidget.SetText(context.format.Render(context.date))
}
