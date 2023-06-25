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

func Test_given_the_current_date_Render_should_produce_a_format_for_the_current_date(t *testing.T) {
	format, err := ByKey(FORMAT_t)

	assert.NoError(t, err)
	assert.Equal(t, format.Key, FORMAT_t)

	var when = time.Now()
	assert.Equal(t, fmt.Sprintf("<t:%d:%s>", when.Unix(), FORMAT_t), format.Render(when))
}
