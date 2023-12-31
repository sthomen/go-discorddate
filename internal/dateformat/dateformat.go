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

var Formats = []DateFormat{
	DateFormat{FORMAT_R, "in 2 hours"},
	DateFormat{FORMAT_D, "November 4, 2023"},
	DateFormat{FORMAT_d, "04/11/2023"},
	DateFormat{FORMAT_T, "11:28:27 AM"},
	DateFormat{FORMAT_t, "11:28 AM"},
	DateFormat{FORMAT_F, "Saturday, November 4, 2023 11:28:27 AM"},
	DateFormat{FORMAT_f, "4 November 2023 11:28"},
}

func ByKey(which string) (DateFormat, error) {
	var result DateFormat
	var err error = errors.New("Invalid format")

	for _, format := range Formats {
		if format.Key == which {
			result = format
			err = nil
			break
		}
	}

	return result, err
}

func ByIndex(index int) (DateFormat, error) {
	var result DateFormat
	var err error = errors.New("Index out of range")

	for i, format := range Formats {
		if i == index {
			result = format
			err = nil
		}
	}

	return result, err
}

func (self *DateFormat)ToIndex() int {
	for index, format := range Formats {
		if format.Key == self.Key {
			return index
		}
	}

	/* NOTREACHED */
	return -1
}

func (self *DateFormat)Render(date time.Time) string {
	return fmt.Sprintf("<t:%d:%s>", date.Unix(), self.Key)
}
