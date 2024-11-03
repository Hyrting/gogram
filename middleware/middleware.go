package middleware

import (
	"errors"
	"log"

	gogram "github.com/hyrting/gogram"
)

// AutoRespond returns a middleware that automatically responds
// to every callback.
func AutoRespond() gogram.MiddlewareFunc {
	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		return func(c gogram.Context) error {
			if c.Callback() != nil {
				defer c.Respond()
			}
			return next(c)
		}
	}
}

// IgnoreVia returns a middleware that ignores all the
// "sent via" messages.
func IgnoreVia() gogram.MiddlewareFunc {
	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		return func(c gogram.Context) error {
			if msg := c.Message(); msg != nil && msg.Via != nil {
				return nil
			}
			return next(c)
		}
	}
}

type RecoverFunc = func(error, gogram.Context)

// Recover returns a middleware that recovers a panic happened in
// the handler.
func Recover(onError ...RecoverFunc) gogram.MiddlewareFunc {
	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		return func(c gogram.Context) error {
			var f RecoverFunc
			if len(onError) > 0 {
				f = onError[0]
			} else if b, ok := c.Bot().(*gogram.Bot); ok {
				f = b.OnError
			} else {
				f = func(err error, _ gogram.Context) {
					log.Println("gogram/middleware/recover:", err)
				}
			}

			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(error); ok {
						f(err, c)
					} else if s, ok := r.(string); ok {
						f(errors.New(s), c)
					}
				}
			}()

			return next(c)
		}
	}
}