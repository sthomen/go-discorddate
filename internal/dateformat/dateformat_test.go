package dateformat

import (
	"fmt"
	"time"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_given_an_invalid_key_ByKey_returns_err(t *testing.T) {
	_, err := ByKey("foo")
	assert.Error(t, err, "An invalid key must return an error")
}

func Test_given_an_empty_key_ByKey_returns_err(t *testing.T) {
	_, err := ByKey("")
	assert.Error(t, err, "An empty key must return an error")
}

func Test_given_a_the_key_FORMAT_t_key_ByKey_returns_t_dateFormat(t *testing.T) {
	result, err := ByKey(FORMAT_t)

	assert.NoError(t, err, "A valid key must not produce an error")
	assert.Equal(t, result.Key, FORMAT_t, "The key '%v' did not match the expected '%v'", result.Key, FORMAT_t)
}

func Test_given_an_invalid_index_ByIndex_returns_err(t *testing.T) {
	_, err := ByIndex(100)
	assert.Error(t, err, "An invalid index must return an error")
}

func Test_given_a_valid_index_ByIndex_returns_the_appropriate_dateFormat(t *testing.T) {
	var index int = 0 // index 0 should be FORMAT_R
	result, err := ByIndex(index)
	assert.NoError(t, err, "A valid index must not return an error")
	assert.Equal(t, result.Key, FORMAT_R, "The index '%v' did not match the expected format '%v'", index, FORMAT_R)
}

func Test_given_a_format_ToIndex_returns_the_expected_index(t *testing.T) {
	for i := 0; i < len(formats); i++ {
		assert.Equal(t, i, formats[i].ToIndex())
	}
}

func Test_given_a_format_that_somehow_is_not_in_formats_ToIndex_returns_negative_one(t *testing.T) {
	format := DateFormat{"X", "Format X"}

	assert.Equal(t, -1, format.ToIndex())
}

func Test_given_the_current_date_Render_should_produce_a_format_for_the_current_date(t *testing.T) {
	format, err := ByKey(FORMAT_t)

	assert.NoError(t, err)
	assert.Equal(t, format.Key, FORMAT_t)

	var when = time.Now()
	assert.Equal(t, fmt.Sprintf("<t:%d:%s>", when.Unix(), FORMAT_t), format.Render(when))
}
