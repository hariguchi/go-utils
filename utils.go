package utils

/*
Copyright 2024 Yoichi Hariguchi

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
“Software”), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

import (
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/sirupsen/logrus"
)

type IPtypes int
type IPindex int

const (
	IPv4  IPtypes = 4
	IPv6  IPtypes = 6
	IPv4i IPindex = 0
	IPv6i IPindex = 1

	Panic = logrus.PanicLevel
	Fatal = logrus.FatalLevel
	Error = logrus.ErrorLevel
	Warn  = logrus.WarnLevel
	Info  = logrus.InfoLevel
	Debug = logrus.DebugLevel
	Trace = logrus.TraceLevel
)

// How to access: TCP[(IPtypes >> 1) & 1]
var (
	TCP = []string{"tcp4", "tcp6"}
	UDP = []string{"udp4", "udp6"}
)

// Regular expressions
var (
	fNameRe = regexp.MustCompile(`^.*/[^/]*\.([^/]*)$`)
	mNameRe = regexp.MustCompile(`^.*/[^/]*\.(\(\*[^/]*\)\.[^/]*)$`)
)

func Eprintf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func SetLogFormat(log *logrus.Logger) {
	// Add file name, line number, and function name to logs
	log.SetReportCaller(true)

	// Time stamp format: UTC
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

func FuncName(full bool) string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		return ""
	}

	fName := runtime.FuncForPC(counter).Name()
	if full {
		return fName
	}
	m := mNameRe.FindAllStringSubmatch(fName, -1)
	if m == nil {
		m = fNameRe.FindAllStringSubmatch(fName, -1)
		if m == nil {
			return ""
		}
		if len(m) < 1 {
			return fName
		}
	}
	if len(m) < 1 {
		return fName
	}
	return m[0][1]
}
