package gui

import (
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/container"

	"git.shangtai.net/staffan/go-discorddate/internal/dateformat"
)

func MainWindow(time time.Time, format dateformat.DateFormat) {
	app := app.New()
	w := app.NewWindow("Discord Date Tool")

	formatWidget := widget.NewSelect(dateformat.FormatsSlice(), func (value string) {
		print(value)
	})

	w.SetContent(container.NewVBox(formatWidget))
	w.ShowAndRun()
}
