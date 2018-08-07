package logs

import (
	"testing"
)


func Test_log(t *testing.T) {
		log := NewLogger(10000)
		log.SetLogger("console", "{}")
		log.SetLogger("file", `{"filename": "test.log"}`)
		log.SetLogger("syslog", `{"level":4}`)

		log.Trace("trace")
		log.Info("info")
		log.Warn("warning")
		log.Debug("debug")
		log.Critical("critical")
}