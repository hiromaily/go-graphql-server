package debug

import (
	"github.com/bookerzzz/grok"
)

func DigIn(value interface{}, options ...grok.Option) {
	grok.Value(value, options...)
}
