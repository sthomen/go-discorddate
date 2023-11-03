package tz

import (
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_given_a_time_in_UTC_AdjustToLocal_changes_the_timezone_without_also_changing_the_time(t *testing.T) {
	when, _ := time.Parse(FORMAT, "2023-11-04 00:00:00")
	adjusted := AdjustToLocal(when)

	assert.Equal(t, when.Format(FORMAT), adjusted.Format(FORMAT))
	assert.Equal(t, adjusted.Location(), time.Now().Location())
}

func Test_given_a_time_in_Local_AdjustToUTC_changes_the_timezone_without_also_changing_the_time(t *testing.T) {
	when, _ := time.ParseInLocation(FORMAT, "2023-11-04 00:00:00", time.Now().Location())
	adjusted := AdjustToUTC(when)

	assert.Equal(t, when.Format(FORMAT), adjusted.Format(FORMAT))
	assert.Equal(t, adjusted.Location(), time.UTC)
}
