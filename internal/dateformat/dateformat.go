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

type DateFormat struct {
	Key  string
	Name string
}

var formats = map[string]DateFormat{
	FORMAT_R: DateFormat{FORMAT_R, "in 2 hours"},
	FORMAT_D: DateFormat{FORMAT_D, "November 4, 2023"},
	FORMAT_d: DateFormat{FORMAT_d, "04/11/2023"},
	FORMAT_T: DateFormat{FORMAT_T, "11:28:27 AM"},
	FORMAT_t: DateFormat{FORMAT_t, "11:28 AM"},
	FORMAT_F: DateFormat{FORMAT_F, "Saturday, November 4, 2023 11:28:27 AM"},
	FORMAT_f: DateFormat{FORMAT_f, "4 November 2023 11:28"},
}

func ByKey(which string) (DateFormat, error) {
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

func FormatsOptions() []string {
	var result = make([]string, 0, len(formats))

	for _, format := range formats {
		result = append(result, format.Name)
	}

	return result
}

func (self *DateFormat) Render(date time.Time) string {
	return fmt.Sprintf("<t:%d:%s>", date.Unix(), self.Key)
}
