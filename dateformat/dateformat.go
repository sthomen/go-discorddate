package dateformat

import (
	"fmt"
	"time"
	"errors"
)

const (
	FORMAT_R = "R"
	FORMAT_D = "D"
	FORMAT_d = "d"
	FORMAT_T = "T"
	FORMAT_t = "t"
	FORMAT_F = "F"
	FORMAT_f = "f"
)

type dateFormat struct {
	Key  string
	Name string
}

var formats = map[string]dateFormat{
	"R": dateFormat{"R", "in 2 hours"},
	"D": dateFormat{"D", "November 4, 2023"},
	"d": dateFormat{"d", "04/11/2023"},
	"T": dateFormat{"T", "11:28:27 AM"},
	"t": dateFormat{"t", "11:28 AM"},
	"F": dateFormat{"F", "Saturday, November 4, 2023 11:28:27 AM"},
	"f": dateFormat{"f", "4 November 2023 11:28"},
}

func Get(which string) (dateFormat, error) {
	result, ok := formats[which]

	var err error = nil

	if !ok {
		err = errors.New("Invalid format: " + which)
	}

	return result, err
}

func Formats() map[string]string {
	var result = make(map[string]string)

	for key,format := range formats {
		result[key] = format.Name
	}

	return result
}

func (self *dateFormat) Render(date time.Time) string {
	return fmt.Sprintf("<t:%d:%s>", date.Unix(), self.Key)
}
