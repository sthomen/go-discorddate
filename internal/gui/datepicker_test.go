package gui

import (
	"time"
	"testing"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/stretchr/testify/assert"
)

func Test_DatePicker_GetDateTime_returns_the_date_set_with_SetDateTime(t *testing.T) {
	// not sure about the defer, the fyne tests do it
	test.NewApp()
	defer test.NewApp()
	
	datePicker := NewDatePicker(func(time.Time) {})
	w := test.NewWindow(datePicker)

	defer w.Close()

	var when = time.Date(1980, time.Month(11), 4, 12, 34, 0, 0, time.UTC)

	datePicker.SetDateTime(when)
	assert.Equal(t, when, datePicker.GetDateTime())
}

func Test_changing_the_date_should_change_the_Time_returned_by_GetDateTime(t *testing.T) {
	defer test.NewApp()
	
	datePicker := NewDatePicker(func(time.Time) {})
	w := test.NewWindow(datePicker)

	defer w.Close()

	// this somehow triggers a lot of setting of the hour and date entry items
	var objects = test.LaidOutObjects(datePicker)
	var facit time.Time

	datePicker.SetDateTime(time.Date(1980, time.Month(11), 4, 12, 34, 0, 0, time.UTC))

	// year
	facit = time.Date(2001, time.Month(11), 4, 12, 34, 0, 0, time.UTC)
	objects[2].(*widget.Select).SetSelected("2001")
	assert.Equal(t, facit, datePicker.GetDateTime())

	// month (XXX translation?)
	facit = time.Date(2001, time.Month(12), 4, 12, 34, 0, 0, time.UTC)
	objects[10].(*widget.Select).SetSelected("December")
	assert.Equal(t, facit, datePicker.GetDateTime())

	// day
	facit = time.Date(2001, time.Month(12), 24, 12, 34, 0, 0, time.UTC)
	objects[18].(*widget.Select).SetSelected("24")
	assert.Equal(t, facit, datePicker.GetDateTime())

	// hour
	facit = time.Date(2001, time.Month(12), 24, 15, 34, 0, 0, time.UTC)
	objects[26].(*widget.Entry).SetText("15")
	// XXX sleep a little so that the UI has time to react
	time.Sleep(50 * time.Millisecond)
	assert.Equal(t, facit, datePicker.GetDateTime())

	// minute
	facit = time.Date(2001, time.Month(12), 24, 15, 45, 0, 0, time.UTC)
	objects[50].(*widget.Entry).SetText("45")
	// XXX sleep a little so that the UI has time to react
	time.Sleep(50 * time.Millisecond)
	assert.Equal(t, facit, datePicker.GetDateTime())
}
