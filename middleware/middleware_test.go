package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gogram "github.com/hyrting/gogram"
)

var b, _ = gogram.NewBot(gogram.Settings{Offline: true})

func TestRecover(t *testing.T) {
	onError := func(err error, c gogram.Context) {
		require.Error(t, err, "recover test")
	}

	h := func(c gogram.Context) error {
		panic("recover test")
	}

	assert.Panics(t, func() {
		h(nil)
	})

	assert.NotPanics(t, func() {
		Recover(onError)(h)(nil)
	})
}
