package gui

import (
	"time"
	"strconv"
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/data/binding"
)

type DatePicker struct {
	widget.BaseWidget

	// change callback
	OnChange     func(time.Time)

	// Internal time state
	when         time.Time

	// internal widgets
	yearWidget   *widget.Select
	monthWidget  *widget.Select
	dayWidget    *widget.Select
	hour         binding.Int
	minute       binding.Int
}

func NewDatePicker(callback func(time.Time)) *DatePicker {
	datePicker := &DatePicker{BaseWidget: widget.BaseWidget{}, OnChange: callback}
	datePicker.ExtendBaseWidget(datePicker)
	return datePicker
}

func (self *DatePicker)SetDateTime(when time.Time) {
	self.when = when
	self.updateAll()
}

func (self *DatePicker)GetDateTime() time.Time {
	return self.when
}

func (self *DatePicker)CreateRenderer() fyne.WidgetRenderer {
	self.ExtendBaseWidget(self)

	self.yearWidget = widget.NewSelect(yearsList(), self.yearChanged)
	self.monthWidget = widget.NewSelect(monthsList(), self.monthChanged)
	self.dayWidget = widget.NewSelect(daysList(self.when), self.dayChanged) 

	self.hour = binding.NewInt()
	self.hour.AddListener(binding.NewDataListener(self.hourChanged))
	self.minute = binding.NewInt()
	self.minute.AddListener(binding.NewDataListener(self.minuteChanged))

	hourWidget := widget.NewEntryWithData(binding.IntToString(self.hour))
	minuteWidget := widget.NewEntryWithData(binding.IntToString(self.minute))

	hourWidget.Validator = func(s string) error {
		i, err := strconv.Atoi(s)

		if err != nil {
			return err
		}

		if len(s) > 2 || i < 0 || i > 23 {
			return errors.New("Hours out of range")
		}

		return nil
	}

	minuteWidget.Validator = func(s string) error {
		i, err := strconv.Atoi(s)

		if err != nil {
			return err
		}

		if len(s) > 2 || i < 0 || i > 59 {
			return errors.New("Minutes out of range")
		}

		return nil
	}

	dateContainer := container.NewHBox(
		self.yearWidget,
		self.monthWidget,
		self.dayWidget,
		hourWidget,
		minuteWidget,
	)

	return widget.NewSimpleRenderer(dateContainer)
}

func (self *DatePicker)yearChanged(year string) {
	parsed, err := strconv.Atoi(year)

	if err == nil {
		self.when = time.Date(
			parsed,
			self.when.Month(),
			self.when.Day(),
			self.when.Hour(),
			self.when.Minute(),
			0,
			0,
			self.when.Location())

		self.OnChange(self.when)
	}
}

func (self *DatePicker)monthChanged(_ string) {
	var month = self.monthWidget.SelectedIndex()

	self.when = time.Date(
		self.when.Year(),
		time.Month(month + 1),
		self.when.Day(),
		self.when.Hour(),
		self.when.Minute(),
		0,
		0,
		self.when.Location())

	self.dayWidget.Options = daysList(self.when)
	self.dayWidget.Refresh()

	self.OnChange(self.when)
}

func (self *DatePicker)dayChanged(day string) {
	parsed, err := strconv.Atoi(day)

	if err == nil {
		self.when = time.Date(
			self.when.Year(),
			self.when.Month(),
			parsed,
			self.when.Hour(),
			self.when.Minute(),
			0,
			0,
			self.when.Location())

		self.OnChange(self.when)
	}
}

func (self *DatePicker)hourChanged() {
	hour, err := self.hour.Get()
	
	if err == nil {
		self.when = time.Date(
			self.when.Year(),
			self.when.Month(),
			self.when.Day(),
			hour,
			self.when.Minute(),
			0,
			0,
			self.when.Location())

		self.OnChange(self.when)
	}
}

func (self *DatePicker)minuteChanged() {
	minute, err := self.minute.Get()

	if err == nil {
		self.when = time.Date(
			self.when.Year(),
			self.when.Month(),
			self.when.Day(),
			self.when.Hour(),
			minute,
			0,
			0,
			self.when.Location())

		self.OnChange(self.when)
	}
}

func (self *DatePicker)updateAll() {
	self.yearWidget.SetSelected(strconv.Itoa(self.when.Year()))
	self.monthWidget.SetSelected(self.when.Month().String())
	self.dayWidget.SetSelected(strconv.Itoa(self.when.Day()))
	
	self.hour.Set(self.when.Hour())
	self.minute.Set(self.when.Minute())

	self.OnChange(self.when)
}

func yearsList() []string {
	var now = time.Now()
	var result []string

	// XXX These dates are a bit arbitrary, the starting year of 1970 stems
	// from the fact that this program really only deals with unix times and
	// the now + 10 seemed reasonable, who would want a date more than 10 years
	// in the future?
	for i := 1970; i <= now.Year() + 10; i++ {
		result = append(result, strconv.Itoa(i))
	}

	return result
}

func monthsList() []string {
	var result []string

	for i := 1; i <= 12; i++ {
		result = append(result, time.Month(i).String())
	}

	return result
}

func daysList(when time.Time) []string {
	cur := time.Date(when.Year(), when.Month(), 1, 12, 0, 0, 0, when.Location())

	var result []string

	for cur.Month() == when.Month() {
		result = append(result, strconv.Itoa(cur.Day()))
		cur = cur.AddDate(0, 0, 1)
	}

	return result
}
