package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// log.Lshortfile 支持显示文件名和代码行号
var (
	errorlog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infolog  = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorlog, infolog}
	mutex    sync.Mutex
)

// 日志打印方法封装暴露
var (
	Error  = errorlog.Println
	Errorf = errorlog.Printf
	Info   = infolog.Println
	Infof  = infolog.Printf
)

// 日志层级 通过控制 Output，来控制日志是否打印。
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

//
func SetLevel(level int) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	if ErrorLevel < level {
		errorlog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infolog.SetOutput(ioutil.Discard)
	}
}
