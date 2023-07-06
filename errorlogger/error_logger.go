package errorlogger

import (
	"github.com/getsentry/sentry-go"
)

func LogError(err error, metadata ...string) {
	event := sentry.NewEvent()
	event.Message = err.Error()

	if len(metadata) > 0 {
		event.Extra["metadata"] = metadata
	}

	// Manually capture the error with the event hint
	sentry.CaptureEvent(event)
}
