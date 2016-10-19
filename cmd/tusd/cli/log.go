package cli

import (
	"log"
	"os"

	"github.com/nagae-memooff/tusd"
)

var stdout = log.New(os.Stdout, "[tusd] ", 0)
var stderr = log.New(os.Stderr, "[tusd] ", 0)

func logEv(eventName string, details ...string) {
	tusd.LogEvent(stderr, eventName, details...)
}
