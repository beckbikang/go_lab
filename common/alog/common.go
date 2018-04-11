package alog

/**

1 日志编写的通用组件

**/
import (
	"fmt"
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

//定义工厂
type LogFactory func() Logger

var logFactories = map[LOG_TYPE]LogFactory{}

//注册方法
