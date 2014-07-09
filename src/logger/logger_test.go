package logger

import (
	. "config"
	"os"
	"testing"
)

func TestErrorln(t *testing.T) {
	Errorln("test errorln func")
	_, err := os.Stat(ERRORPATH)
	if os.IsExist(err) {
		return
	} else {
		t.Fail()
	}
}
