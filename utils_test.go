package utils

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type bag struct {
	nItems int
}

func TestLogFormat(t *testing.T) {
	fName := FuncName(false)
	Eprintf("=== %s ===\n", fName)

	log := logrus.New()
	SetLogFormat(log)
	log.Infof("Testing...")
}

func TestL4Protos(t *testing.T) {
	fName := FuncName(false)
	Eprintf("=== %s ===\n", fName)

	tcp := TCP[(IPv4>>1)&1]
	if tcp == "tcp4" {
		Eprintf(`TCP[(IPv4>>1)&1] == "tcp4"` + "\n")
	} else {
		t.Fatalf(`Result %s must be "'tcp4"`+"\n", tcp)
	}
	tcp = TCP[(IPv6>>1)&1]
	if tcp == "tcp6" {
		Eprintf(`TCP[(IPv6>>1)&1] == "tcp6"` + "\n")
	} else {
		t.Fatalf(`Result %s must be "'tcp6"`+"\n", tcp)
	}
	udp := UDP[(IPv4>>1)&1]
	if udp == "udp4" {
		Eprintf(`UDP[(IPv4>>1)&1] == "udp4"` + "\n")
	} else {
		t.Fatalf(`Result %s must be "'udp4"`+"\n", udp)
	}
	udp = UDP[(IPv6>>1)&1]
	if udp == "udp6" {
		Eprintf(`UDP[(IPv4>>1)&1] == "udp6"` + "\n")
	} else {
		t.Fatalf(`Result %s must be "'udp6"`+"\n", udp)
	}
}

func TestFuncName(t *testing.T) {
	fName := FuncName(false)
	Eprintf("=== %s ===\n", fName)

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
