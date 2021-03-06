package air

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"
	"time"
)

// logger is an active logging object that generates lines of output.
type logger struct {
	template *template.Template
	once     *sync.Once
}

// theLogger is the singleton of the `logger`.
var theLogger = &logger{
	template: template.New("logger"),
	once:     &sync.Once{},
}

// log logs the v at the level.
func (l *logger) log(level string, v ...interface{}) {
	if !LoggerEnabled {
		return
	}
	l.once.Do(func() {
		template.Must(l.template.Parse(LoggerFormat))
	})
	buf := bytes.Buffer{}
	l.template.Execute(&buf, map[string]interface{}{
		"AppName": AppName,
		"Time":    time.Now().UTC().Format(time.RFC3339),
		"Level":   level,
		"Message": fmt.Sprint(v...),
	})
	LoggerOutput.Write(buf.Bytes())
}
