package utils

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type bag struct {
	nItems int
}

func TestLogFormat(t *testing.T) {
	log := logrus.New()
	SetLogFormat(log)
	log.Infof("Testing...")
}

func TestFuncName(t *testing.T) {
	fName := FuncName(false)
	Eprintf("%s\n", fName)
	if fName != "TestFuncName" {
		t.Fatalf("result %s != TestFuncName\n", fName)
	}
	Eprintf("%s\n", FuncName(true))

	b := &bag{nItems: 0}
	fName = b.ShowMethodName(false)
	Eprintf("%s\n", fName)
	if fName != "(*bag).ShowMethodName" {
		t.Fatalf("result %s != (*bag).ShowMethodName\n", fName)
	}
	Eprintf("%s\n", b.ShowMethodName(true))
}

func (p *bag) ShowMethodName(full bool) string {
	return FuncName(full)
}
