package fly

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

// Duration represents the elapsed time between two instants as an int64 nanosecond count.
type Duration struct {
	d time.Duration
}

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
	Week                      = 7 * Day
)

func newDuration(d time.Duration) *Duration {
	return &Duration{d}
}

// ParseDuration parses a duration string
func ParseDuration(s string) (*Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return nil, err
	}

	return newDuration(d), nil
}

// Since returns the time elapsed since t.
func Since(f Fly) *Duration {
	return newDuration(time.Since(f.t))
}

// Hours returns the duration as a floating point number of hours
func (d *Duration) Hours() float64 {
	return d.d.Hours()
}

// Hour returns the hour elapsed
func (d *Duration) Hour() int {
	h := d.d % Hour
	h = h / Hour
	return int(h)
}

// Minutes returns the duration as a floating point number of minutes
func (d *Duration) Minutes() float64 {
	return d.d.Minutes()
}

// Seconds returns the duration as a floating point number of seconds
func (d *Duration) Seconds() float64 {
	return d.d.Seconds()
}

// Nanoseconds returns the duration as an integer nanosecond count
func (d *Duration) Nanoseconds() int64 {
	return d.d.Nanoseconds()
}

// String returns a string representing the duration in the form "72h3m0.5s".
func (d *Duration) String() string {
	return d.d.String()
}

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

// Millisecond returns current time milliseconds value
func (f *Fly) Millisecond() int {
	return f.t.Nanosecond() / 1000 / 1000
}

// Microsecond returns current time milliseconds value
func (f *Fly) Microsecond() int {
	return (f.t.Nanosecond() / 1000) % 1000
}

// Nanosecond returns current time milliseconds value
// Note: Nanosecond is in the range [0, 999999999], which includes seconds and microseconds.
func (f *Fly) Nanosecond() int {
	return f.t.Nanosecond()
}

// pastHalf checks if the time passes half of the duration.
// If d is 1 hour, it returns if it's somewhere past 30 minutes.
func (f *Fly) pastHalf(d time.Duration) bool {
	switch d {
	case time.Duration(time.Hour):
		return f.t.Minute() > 30
	case time.Duration(time.Minute):
		return f.t.Second() > 30
	case time.Duration(time.Second):
		return f.Millisecond() > 500
	case time.Duration(time.Millisecond):
		return f.Microsecond() > 500
	case time.Duration(time.Microsecond):
		return f.t.Nanosecond() > 500
	default:
		// TODO: handle unformatted duration
		return false
	}
}

// Floor returns the result of time floor of certain unit.
// Parameter should be hour, minute, second, microsecond,
// If unit is negative, Floor returns time unchanged
func (f *Fly) Floor(name string) (*Fly, error) {
	// TODO(cizixs): format duration to standard unit
	d, err := time.ParseDuration(name)
	if err != nil {
		return nil, err
	}
	newT := f.t.Round(d)
	if f.pastHalf(d) {
		newT = newT.Add(-d)
	}
	return New(newT), nil
}

// Ceil returns the result of time ceil of certain unit.
// Parameter should be hour, minute, second, microsecond,
// If unit is negative, Ceil returns time unchanged
func (f *Fly) Ceil(name string) (*Fly, error) {
	// TODO(cizixs): format duration to standard unit
	d, err := time.ParseDuration(name)
	if err != nil {
		return nil, err
	}
	newT := f.t.Round(d)
	if !f.pastHalf(d) {
		newT = newT.Add(time.Duration(d))
	}
	return New(newT), nil
}

// UTCNow returns right now based on utc time zone
func UTCNow() *Fly {
	return &Fly{time.Now().UTC()}
}
