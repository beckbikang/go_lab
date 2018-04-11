package alog

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

/**

1 日志的级别
	trace,info,warn,error,fatal
2 日志的类型
	各种基本级别的日志操作类

**/
type (
	LOG_LEVEL int
	LOG_TYPE  string
)

var (
	TRACE = LOG_LEVEL(0)
	INFO  = LOG_LEVEL(1)
	WARN  = LOG_LEVEL(2)
	ERROR = LOG_LEVEL(3)
	FATAL = LOG_LEVEL(4)
)

//判断日志级别是否可用
func (level LOG_LEVEL) IsValid() bool {
	return level >= TRACE && level <= FATAL
}

//打印日志
func (level LOG_LEVEL) String() string {
	switch level {
	case TRACE:
		return "[TRACE] "
	case INFO:
		return "[ INFO] "
	case WARN:
		return "[ WARN] "
	case ERROR:
		return "[ERROR] "
	case FATAL:
		return "[FATAL] "
	}
}

//日志数据
type LogData struct {
	Level LOG_LEVEL
	Data  string
}

//最基本的打印日志数据
func WriteLog(level LOG_LEVEL, skip int,
	format string, v ...interface{}) {

	logData := &LogData{
		Level: level,
	}

	if level >= FATAL && skip > 0 {
		pc, file, line, ok = runtime.Caller(skip)
		if ok {
			//返回一个表示调用栈标识符pc对应的调用栈的*Func；
			fn := runtime.FuncForPC(pc)
			var funcName string
			if fn == nil {
				funcName = "?()"
			} else {
				//返回path文件扩展名
				funcName = strings.TrimLeft(filepath.Ext(file.Name()), ".") + "()"
			}
			if len(file) > 20 {
				file = "..." + file[len(file)-20:]
			}
			logData.Data = level.String() + fmt.Sprintf("[%s:%d %s]", file,
				line, funcName) + fmt.Sprintf(format, v...)
		}
	}
	if len(logData.Data) == 0 {
		logData.Data = level.String() + fmt.Sprintf(format, v...)
	}
	//todo add something
}

func Trace(format string, v ...interface{}) {
	WriteLog(TRACE, 0, format, v...)
}

func Info(format string, v ...interface{}) {
	WriteLog(INFO, 0, format, v...)
}

func Warn(format string, v ...interface{}) {
	WriteLog(WARN, 0, format, v...)
}

func Error(skip int, format string, v ...interface{}) {
	WriteLog(TRACE, skip, format, v...)
}

func Fatal(skip int, format string, v ...interface{}) {
	WriteLog(TRACE, skip, format, v...)
	Shutdown()
	os.Exit(1)
}

func Shutdown() {

}
