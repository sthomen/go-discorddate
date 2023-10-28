package gui

import (
	"fmt"
	"time"

	"github.com/libui-ng/golang-ui"
	_ "github.com/libui-ng/golang-ui/winmanifest"
	"golang.design/x/clipboard"

	"git.shangtai.net/staffan/go-discorddate/internal/dateformat"
)

// gui context, used internally for widget references
var context struct {
	defaults struct {
		t time.Time
		f dateformat.DateFormat
	}
	date   *ui.DateTimePicker
	format *ui.Combobox
	label  *ui.Label
}

// Create the main window and run the UI loop
func MainWindow(when time.Time, format dateformat.DateFormat) {
	var err error

	err = clipboard.Init()

	if err != nil {
		panic(err)
	}

	context.defaults.t = when
	context.defaults.f = format

	err = ui.Main(setupUi)

	if err != nil {
		panic(err)
	}
}

// libui UI setup callback
func setupUi() {
	mainWindow := ui.NewWindow("Discord Date Tool", 0, 0, false)
	mainWindow.SetMargined(true)

	mainWindow.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	ui.OnShouldQuit(func() bool {
		mainWindow.Destroy()
		return true
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)

	context.date    = ui.NewDateTimePicker()
	context.format  = ui.NewCombobox()
	button         := ui.NewButton("Copy to clipboard")
	context.label   = ui.NewLabel("")

	for _, format := range dateformat.Formats {
		context.format.Append(format.Name)
	}

	context.date.OnChanged(func(_ *ui.DateTimePicker) {
		context.label.SetText(makeString())
	})

	context.format.OnSelected(func(_ *ui.Combobox) {
		context.label.SetText(makeString())
	})

	defaultValues()

	button.OnClicked(func(_ *ui.Button) {
		str := makeString()

		if str != "" {
			clipboard.Write(clipboard.FmtText, []byte(str))
			context.label.SetText(fmt.Sprintf("Copied: %s", str))
		} else {
			context.label.SetText("Insufficient parameter")
		}
	})

	vbox.Append(context.date, false)
	vbox.Append(context.format, false)
	vbox.Append(button, false)
	vbox.Append(context.label, false)
	mainWindow.SetChild(vbox)

	mainWindow.Show()
}

// set the UI to the default values as passed in by the command line options
func defaultValues() {
	context.date.SetTime(context.defaults.t)
	context.format.SetSelected(context.defaults.f.ToIndex())
	context.label.SetText(makeString())
}

// Utility for pulling the selected time and format from the widgets and rendering
// the date string.
func makeString() string {
	when := context.date.Time()
	format, err := dateformat.ByIndex(context.format.Selected())

	if err != nil {
		return "" /* NOTREACHED */
	}

	return format.Render(when)
}
