package fly

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// Fly is the main data struct
type Fly struct {
	t time.Time
}

// New creates and returns a Fly struct from `time.Time` instance
func New(t time.Time) *Fly {
	return &Fly{t: t}
}

// Now creates an instance with time set to now()
func Now() *Fly {
	return New(time.Now())
}

// Humanize return the human readable date string
func (f *Fly) Humanize() string {
	return humanize.Time(f.t)
}

// String returns the formatted output time
func (f *Fly) String() string {
	return f.t.String()
}

// Zone return the timezone name, and its offset from UTC
func (f *Fly) Zone() (string, int) {
	return f.t.Zone()
}

// Add moves the time forward by certain mount of duration.
// Parameter can be either `time.Duration` instance or time string
// such as "300ms", "-1.5h", "2h 45m"(spaces are handled carefully),
// each with a decimal number and a time unit.
// Valid time units are "ns", "us", "ms", "s", "m", "h".
//
// Note that you can move time backward by setting the value to negative, like
// `time.Duration(-5 * time.Hour)`, or `-5h`.
// If parameter can not be parsed as duration, an error will be returned
func (f *Fly) Add(d interface{}) (*Fly, error) {
	offset := time.Duration(0)
	var err error
	switch d.(type) {
	case time.Duration:
		offset = d.(time.Duration)
	case string:
		value := d.(string)
		value = strings.Replace(value, " ", "", -1)
		offset, err = time.ParseDuration(value)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Unknow duration instance.")
	}
	return &Fly{f.t.Add(offset)}, nil
}

// To returns the time of the location identified by given name.
// If name is "" or "UTC", To returns UTC time. If the name is "Local",
// To returns local time.
// Otherwsie, To returns the time at location corresponding to IANA Time Zone,
// such as "America/New_York", or "Asia/Shanghai".
func (f *Fly) To(name string) (*Fly, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}

	return New(f.t.In(loc)), nil
}

// UTCNow returns right now based on utc time zone
func UTCNow() *Fly {
	return &Fly{time.Now().UTC()}
}
