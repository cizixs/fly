package fly_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cizixs/fly"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	f := fly.New()
	assert.Equal("now", f.Humanize(), "New created instance should output now")
}
