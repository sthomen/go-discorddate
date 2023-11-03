package tz

// Golang's time package does not allow you to change the timezone without also
// changing the value of the time (i.e. golang treats the time as a fixed
// point). While this is useful, it is quite maddening that changing the
// timezone without adjusting the time CANNOT be done without something like
// this.

import "time"

const FORMAT = "2006-01-02 15:04:05"

// Adjust the timezone of the given Time object to UTC without also changing
// the value of the time
func AdjustToUTC(when time.Time) time.Time {
	result, err := time.Parse(FORMAT, when.Format(FORMAT))

	if err != nil {
		panic("Internal error in tz.AdjustToUTC")
	}

	return result
}

// Adjust the timezone of the given Time object to the local timezone without
// also changing the value of the time
func AdjustToLocal(when time.Time) time.Time {
	result, err := time.ParseInLocation(FORMAT, when.Format(FORMAT), time.Now().Location())

	if err != nil {
		panic("Internal error in tz.AdjustToLocal")
	}

	return result
}
