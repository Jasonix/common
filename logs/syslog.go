//build !windows

package logs

import (
		"encoding/json"
								"sync"
	"time"
	"log/syslog"
)

type sysLogWriter struct {
	sync.RWMutex // write log order by order and  atomic incr maxLinesCurLines and maxSizeCurSize
	sysWriter *syslog.Writer

	Level int `json:"level"`

	Perm string `json:"perm"`
}

// newFileWriter create a FileLogWriter returning as LoggerInterface.
func newSysWriter() Logger {
	w := &sysLogWriter{
		Level:      LevelTrace,
		Perm:       "default",
	}
	return w
}

func (w *sysLogWriter) Init(jsonConfig string) error {
	err := json.Unmarshal([]byte(jsonConfig), w)
	if err != nil {
		return err
	}
	w.sysWriter, err = syslog.New(syslog.Priority(w.Level), "")
	return err
}

func (w *sysLogWriter) WriteMsg(when time.Time, msg string, level int) error {
	if level > w.Level {
		return nil
	}
	h, _ := formatTimeHeader(when)
	msg = string(h) + msg + "\n"

	w.Lock()
	_, err := w.sysWriter.Write([]byte(msg))
	w.Unlock()
	return err
}

func (w *sysLogWriter) Destroy() {
	w.sysWriter.Close()
}

func (w *sysLogWriter) Flush() {
}

func init() {
	Register(AdapterSyslog, newSysWriter)
}
