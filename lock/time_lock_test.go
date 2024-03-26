package lock

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestManualUnlock(t *testing.T) {
	tl := NewTimedLock(5 * time.Second)
	tl.Lock()
	err := tl.TryLock()
	assert.NotNil(t, err)
	tl.Unlock()
	err = tl.TryLock()
	assert.Nil(t, err)
}

func TestTimeout(t *testing.T) {
	tl := NewTimedLock(2 * time.Second)
	tl.Lock()
	time.Sleep(2 * time.Second)
	err := tl.TryLock()
	assert.Nil(t, err)
}
