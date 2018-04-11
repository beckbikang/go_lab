package alog

/**

1 日志编写的通用组件

**/
import (
	"fmt"
)

const (
	FACTORY_NIL      = "log factory is nil"
	FACTORY_EXIST    = "log factory has existed"
	ERROR_CHAN_COUNT = 5

	ERROR_MSG_FORMAT = "unable to write logdata:%v\n"
)

//通用的日志接口
type Logger interface {
	Level() LOG_LEVEL
	Init(interface{}) error //初始化日志
	ExchangeChans(chan<- error) chan *LogData
	Start()   //开始
	Destroy() //摧毁
}

//所有日志都需要的通用的部分
type LogCommon struct {
	level       LOG_LEVEL
	logDataChan chan *LogData
	quiteChan   chan struct{}
	errorChan   chan<- error
}

//获取日志的类
type LogFactory func() Logger

//定义日志工厂
var logFactories = map[LOG_TYPE]LogFactory{}

/**
注册各种类型的日志工厂
*/
func Register(logType LOG_TYPE, logFactory LogFactory) {
	if logFactory == nil {
		panic(FACTORY_NIL)
	}
	if logFactories[logType] != nil {
		panic(FACTORY_EXIST)
	}
	logFactories[logType] = logFactory
}

/**
对日志处理器的包装
**/
type logParser struct {
	Logger
	logType     LOG_TYPE      //日志的类型
	logDataChan chan *LogData //消息管道
}

//定义基本常量
var (
	logParsers []*logParser
	errorChan  = make(chan error, ERROR_CHAN_COUNT)
	quitChan   = make(chan struct{})
)

//初始化
func init() {
	go func() {
		for {
			select {
			case err := <-errorChan:
				fmt.Printf(ERROR_MSG_FORMAT, err)
			case quitChan := <-quitChan:
				return
			}
		}
	}()
}

/**
新建一个日志处理器，放入到parsers里面
**/
func NewLogger(logType LOG_TYPE, logConfig interface{}) error {
	fac, ok := logFactories[logType]
	if !ok {
		return fmt.Errorf("%s log type not exist in factories", logType)
	}

	alog := fac()
	if err := alog.Init(logConfig); err != nil {
		return err
	}
	logDataChan := alog.ExchangeChans(errorChan)

	founded := false
	for i := range logParsers {
		if logParsers[i].logType == logType {
			founded = true
			logParsers[i].Destroy()
			logParsers[i].Logger = alog
			logParsers[i].logDataChan = logDataChan
			break
		}

	}
	if !founded {
		logParsers = append(
			logParsers,
			&logParser{
				Logger:      alog,
				logType:     logType,
				logDataChan: logDataChan,
			},
		)
	}

	go alog.Start()
	return nil
}

func DeleteLogger(logType LOG_TYPE) {
	index := -1
	for i := range logParsers {
		if logParsers[i].logType == logType {
			index = i
			logParsers[i].Destroy()
		}
	}
	if index >= 0 {
		newParsers := make([]*logParser, len(logParsers)-1)
		copy(newParsers, logParsers[:index])
		copy(newParsers, logParsers[index+1:])
		logParsers = newParsers
	}
}
