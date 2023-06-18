package dateformat

import (
	"fmt"
	"time"
)

const (
	FORMAT_R = iota
	FORMAT_D
	FORMAT_d
	FORMAT_T
	FORMAT_t
	FORMAT_F
	FORMAT_f
)

type dateFormat struct {
	Key  string
	Name string
}

var formats = map[int]dateFormat{
	0: dateFormat{"R", "in 2 hours"},
	1: dateFormat{"D", "November 4, 2023"},
	2: dateFormat{"d", "04/11/2023"},
	3: dateFormat{"T", "11:28:27 AM"},
	4: dateFormat{"t", "11:28 AM"},
	5: dateFormat{"F", "Saturday, November 4, 2023 11:28:27 AM"},
	6: dateFormat{"f", "4 November 2023 11:28"},
}

func Get(which int) dateFormat {
	return formats[which]
}

func (self *dateFormat) Render(date time.Time) string {
	return fmt.Sprintf("<t:%d:%s>", date.Unix(), self.Key)
}
