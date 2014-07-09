package logger

import (
	. "config"
	"os"
	"testing"
)

func TestErrorln(t *testing.T) {
	Errorln("Test Errorln")
	_, err := os.Stat(ERRORPATH)
	if os.IsNotExist(err) {
		t.Fail()
	} else {
		return
	}
}

func TestDebugln(t *testing.T) {
	Debugln("Test Debugln")
	_, err := os.Stat(DEBUGPATH)
	if os.IsNotExist(err) {
		t.Fail()
	} else {
		return
	}
}
