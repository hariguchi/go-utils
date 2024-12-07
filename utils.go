package utils

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
)

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
