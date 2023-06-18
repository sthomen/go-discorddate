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
	FORMAT_R: dateFormat{FORMAT_R, "in 2 hours"},
	FORMAT_D: dateFormat{FORMAT_D, "November 4, 2023"},
	FORMAT_d: dateFormat{FORMAT_d, "04/11/2023"},
	FORMAT_T: dateFormat{FORMAT_T, "11:28:27 AM"},
	FORMAT_t: dateFormat{FORMAT_t, "11:28 AM"},
	FORMAT_F: dateFormat{FORMAT_F, "Saturday, November 4, 2023 11:28:27 AM"},
	FORMAT_f: dateFormat{FORMAT_f, "4 November 2023 11:28"},
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
