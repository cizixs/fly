package fly_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cizixs/fly"
)

func TestNow(t *testing.T) {
	assert := assert.New(t)
	f := fly.Now()
	assert.Equal("now", f.Humanize(), "New created instance should output now")
}

func TestUTC(t *testing.T) {
	assert := assert.New(t)

	f := fly.UTCNow()
	name, offset := f.Zone()
	assert.Equal("UTC", name, "UTC time zone should be `UTC`.")
	assert.Equal(0, offset, "UTC offset should be zero.")
}

func TestHumanize(t *testing.T) {
	assert := assert.New(t)

	f := fly.Now()

	// Become a hour later, add 2 minutes to make sure it's more than one hour in the next statement
	f, _ = f.Add(time.Duration(time.Hour + 2*time.Minute))
	assert.Equal("1 hour from now", f.Humanize(), "Date should be one hour from now.")
}

func TestAddDuration(t *testing.T) {
	assert := assert.New(t)

	// 2016-12-03 22:15:35 +0000 UTC
	td := time.Date(2016, time.December, 3, 22, 15, 35, 0, time.UTC)
	f := fly.New(td)

	cases := []struct {
		duration interface{} // duration to be added
		err      bool        // if this action will cause error
		expected string      // expected value
	}{
		{time.Duration(-2 * time.Hour), false, "2016-12-03 20:15:35 +0000 UTC"},
		{time.Duration(2 * time.Hour), false, "2016-12-04 00:15:35 +0000 UTC"},
		{time.Duration(3 * time.Minute), false, "2016-12-03 22:18:35 +0000 UTC"},
		{time.Duration(42 * time.Second), false, "2016-12-03 22:16:17 +0000 UTC"},
		{time.Duration(813 * time.Millisecond), false, "2016-12-03 22:15:35.813 +0000 UTC"},
		{"0h", false, "2016-12-03 22:15:35 +0000 UTC"},
		{"-1.5h", false, "2016-12-03 20:45:35 +0000 UTC"},
		{"2h2m", false, "2016-12-04 00:17:35 +0000 UTC"},
		{"300ms", false, "2016-12-03 22:15:35.3 +0000 UTC"},
		{"2h 2m", false, "2016-12-04 00:17:35 +0000 UTC"},
		{"3h ", false, "2016-12-04 01:15:35 +0000 UTC"},
		{"3hours", true, ""},
		{"cizixs", true, ""},
		{34, true, "2016-12-03 22:16:35 +0000 UTC"},
	}

	for _, tt := range cases {
		f, err := f.Add(tt.duration)
		if tt.err {
			assert.Error(err, "Add time duration [%v] should cause error.", tt.duration)
		} else {
			assert.NoError(err, "Add duration should not cause error.")
			assert.Equal(tt.expected, f.String())
		}
	}
}
