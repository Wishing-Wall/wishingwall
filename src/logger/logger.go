package logger

import (
	. "config"
	//"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	os.Mkdir(LOGPATH, 0777)
}

type logger struct {
	*log.Logger
}

func New(out io.Writer) *logger {

	return &logger{
		Logger: log.New(out, "", log.LstdFlags),
	}
}

func Errorln(args ...interface{}) {
	file, err := os.OpenFile(ERRORPATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		log.Println("create errorlog file failed")
		return
	}
	defer file.Close()
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
	}
	New(file).Println(args...)

}
func Debugln(args ...interface{}) {
	file, err := os.OpenFile(DEBUGPATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		return
	}
	defer file.Close()
	_, callerFile, line, ok := runtime.Caller(1)
	if ok {
		args = append([]interface{}{"[", filepath.Base(callerFile), "]", line}, args...)
	}
	New(file).Println(args...)

}
