package fly

import (
	"time"

	"github.com/dustin/go-humanize"
)

// Fly is the main data struct
type Fly struct {
	time.Time
}

// New creates an instance with time set to now()
func New() *Fly {
	now := time.Now()
	return &Fly{
		now,
	}
}

// Humanize return the human readable date string
func (f *Fly) Humanize() string {
	return humanize.Time(f.Time)
}

// UTCNow returns right now based on utc time zone
func UTCNow() *Fly {
	return &Fly{time.Now().UTC()}
}
