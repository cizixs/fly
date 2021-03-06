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

func TestTimeZone(t *testing.T) {
	assert := assert.New(t)

	// 2016-12-03 22:15:35 +0000 UTC
	td := time.Date(2016, time.December, 3, 22, 15, 35, 0, time.UTC)
	f := fly.New(td)

	cases := []struct {
		timezone string
		err      bool
		expected string
	}{
		{"Asia/Shanghai", false, "2016-12-04 06:15:35 +0800 CST"},
		{"America/New_York", false, "2016-12-03 17:15:35 -0500 EST"},
		{"", false, "2016-12-03 22:15:35 +0000 UTC"},
		{"UTC", false, "2016-12-03 22:15:35 +0000 UTC"},
		{"nowhere", true, ""},
	}
	for _, tt := range cases {
		lt, err := f.To(tt.timezone)
		if tt.err {
			assert.Error(err, "Expect error on %v", tt.timezone)
		} else {
			assert.NoError(err, "Convert time to different timezone should cause no error.")
			assert.Equal(tt.expected, lt.String())
		}
	}
}

func TestMillisecond(t *testing.T) {
	assert := assert.New(t)

	// 2016-12-03 22:15:35 +0000 UTC
	td := time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC)
	f := fly.New(td)
	assert.Equal(123, f.Millisecond())
	assert.Equal(456, f.Microsecond())
	assert.Equal(123456789, f.Nanosecond())
}

func TestFloor(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		date     time.Time
		unit     string
		expected string
	}{
		{time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC), "1h", "2016-12-03 22:00:00 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 45, 35, 123456789, time.UTC), "1h", "2016-12-03 22:00:00 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC), "1m", "2016-12-03 22:15:00 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 13, 123456789, time.UTC), "1m", "2016-12-03 22:15:00 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC), "1s", "2016-12-03 22:15:35 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 35, 823456789, time.UTC), "1s", "2016-12-03 22:15:35 +0000 UTC"},
	}
	// 2016-12-03 22:15:35 +0000 UTC

	for _, tc := range cases {
		f := fly.New(tc.date)
		f, _ = f.Floor(tc.unit)
		assert.Equal(tc.expected, f.String())
	}
}

func TestCeil(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		date     time.Time
		unit     string
		expected string
	}{
		{time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC), "1h", "2016-12-03 22:59:59 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 45, 35, 123456789, time.UTC), "1h", "2016-12-03 22:59:59 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC), "1m", "2016-12-03 22:15:59 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 13, 123456789, time.UTC), "1m", "2016-12-03 22:15:59 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 35, 123456789, time.UTC), "1s", "2016-12-03 22:15:35 +0000 UTC"},
		{time.Date(2016, time.December, 3, 22, 15, 35, 823456789, time.UTC), "1s", "2016-12-03 22:15:35 +0000 UTC"},
	}
	// 2016-12-03 22:15:35 +0000 UTC

	for _, tc := range cases {
		f := fly.New(tc.date)
		f, _ = f.Ceil(tc.unit)
		assert.Equal(tc.expected, f.String())
	}
}
