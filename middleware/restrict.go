package middleware

import gogram "github.com/hyrting/gogram"

// RestrictConfig defines config for Restrict middleware.
type RestrictConfig struct {
	// Chats is a list of chats that are going to be affected
	// by either In or Out function.
	Chats []int64

	// In defines a function that will be called if the chat
	// of an update will be found in the Chats list.
	In gogram.HandlerFunc

	// Out defines a function that will be called if the chat
	// of an update will NOT be found in the Chats list.
	Out gogram.HandlerFunc
}

// Restrict returns a middleware that handles a list of provided
// chats with the logic defined by In and Out functions.
// If the chat is found in the Chats field, In function will be called,
// otherwise Out function will be called.
func Restrict(v RestrictConfig) gogram.MiddlewareFunc {
	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		if v.In == nil {
			v.In = next
		}
		if v.Out == nil {
			v.Out = next
		}
		return func(c gogram.Context) error {
			for _, chat := range v.Chats {
				if chat == c.Sender().ID {
					return v.In(c)
				}
			}
			return v.Out(c)
		}
	}
}

// Blacklist returns a middleware that skips the update for users
// specified in the chats field.
func Blacklist(chats ...int64) gogram.MiddlewareFunc {
	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		return Restrict(RestrictConfig{
			Chats: chats,
			Out:   next,
			In:    func(c gogram.Context) error { return nil },
		})(next)
	}
}

// Whitelist returns a middleware that skips the update for users
// NOT specified in the chats field.
func Whitelist(chats ...int64) gogram.MiddlewareFunc {
	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		return Restrict(RestrictConfig{
			Chats: chats,
			In:    next,
			Out:   func(c gogram.Context) error { return nil },
		})(next)
	}
}
