package middleware

import (
	"log"

	"github.com/goccy/go-json"

	gogram "github.com/hyrting/gogram"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger(logger ...*log.Logger) gogram.MiddlewareFunc {
	var l *log.Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = log.Default()
	}

	return func(next gogram.HandlerFunc) gogram.HandlerFunc {
		return func(c gogram.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			l.Println(string(data))
			return next(c)
		}
	}
}
