package binding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type simpleItem struct {
	base
}

func TestBase_AddListener(t *testing.T) {
	data := &simpleItem{}
	assert.Equal(t, 0, len(data.listeners))

	called := false
	fn := NewNotifyFunction(func(DataItem) {
		called = true
	})
	data.AddListener(fn)
	assert.Equal(t, 1, len(data.listeners))

	data.trigger(data)
	assert.True(t, called)
}

func TestBase_RemoveListener(t *testing.T) {
	called := false
	fn := NewNotifyFunction(func(DataItem) {
		called = true
	})
	data := &simpleItem{}
	data.listeners = []DataItemListener{fn}

	assert.Equal(t, 1, len(data.listeners))
	data.RemoveListener(fn)
	assert.Equal(t, 0, len(data.listeners))

	data.trigger(data)
	assert.False(t, called)
}

func TestNewNotifyFunction(t *testing.T) {
	called := false
	fn := NewNotifyFunction(func(DataItem) {
		called = true
	})

	fn.DataChanged(nil)
	assert.True(t, called)
}
