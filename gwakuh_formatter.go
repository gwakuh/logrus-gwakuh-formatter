package logrus

import (
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	pid int
)

var logLevel map[string]string = map[string]string{
	"info":    "LV 1",
	"warning": "LV 2",
	"error":   "LV 3",
	"debug":   "LV 4",
}

const (
	defaultLogFormat       = "%time%, %level%, %pid%, %msg%"
	defaultTimestampFormat = "20060102150405"
)

type Formatter struct {
	TimestampFormat string
	LogFormat       string
}

func init() {
	pid = os.Getpid()
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	out := f.LogFormat
	if out == "" {
		out = defaultLogFormat
	}

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	level, ok := logLevel[entry.Level.String()]
	if ok == false {
		level = entry.Level.String()
	}

	out = strings.Replace(out, "%time%", entry.Time.Format(timestampFormat), 1)
	out = strings.Replace(out, "%level%", level, 1)
	out = strings.Replace(out, "%pid%", strconv.Itoa(pid), 1)
	out = strings.Replace(out, "%msg%", entry.Message, 1)

	for key, value := range entry.Data {
		switch v := value.(type) {
		case string:
			out = strings.Replace(out, "%"+key+"%", v, 1)
		case int:
			out = strings.Replace(out, "%"+key+"%", strconv.Itoa(v), 1)
		case bool:
			out = strings.Replace(out, "%"+key+"%", strconv.FormatBool(v), 1)
		}
	}

	return []byte(out), nil
}
