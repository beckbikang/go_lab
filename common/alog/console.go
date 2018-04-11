package alog

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

const (
	CONSOLE            LOG_TYPE = "CONSOLE"
	CONSOLE_CONFIG_ERR string   = "ConsoleConfig Error"
)

var consoleColors = map[LOG_LEVEL]func(a ...interface{}) string{
	TRACE: color.New(color.FgBlue).SprintFunc(),   // Trace
	INFO:  color.New(color.FgGreen).SprintFunc(),  // Info
	WARN:  color.New(color.FgYellow).SprintFunc(), // Warn
	ERROR: color.New(color.FgRed).SprintFunc(),    // Error
	FATAL: color.New(color.FgHiRed).SprintFunc(),  // Fatal
}

type ConsoleConfig struct {
	Level    LOG_LEVEL
	ChanSize int64
}

type console struct {
	LogCommon
	*log.Logger
}

func newConsole() Logger {
	return &console{
		Logger: log.New(color.Output, "", log.Ldate|log.Ltime),
		LogCommon: LogCommon{
			quiteChan: make(chan struct{}),
		},
	}
}

func (c *console) Level() LOG_LEVEL {
	return c.level
}
func (c *console) Init(cfg interface{}) error {
	config, ok := cfg.(ConsoleConfig)
	if !ok {
		return fmt.Errorf(CONSOLE_CONFIG_ERR)
	}
	if !(config.Level).IsValid() {
		return fmt.Errorf(LEVER_INVALID_ERR)
	}
	c.level = config.Level
	c.logDataChan = make(chan *LogData, config.ChanSize)
	return nil
}

//返回当前的logDataChan
func (c *console) ExchangeChans(errChan chan<- error) chan *LogData {
	c.errorChan = errChan
	return c.logDataChan
}

func (c *console) writelog(logData *LogData) {
	c.Logger.Print(consoleColors[logData.Level](logData.Data))
}

//开始运行
func (c *console) Start() {
LOOP:
	for {
		select {
		case logData := <-c.logDataChan:
			c.writelog(logData)
		case <-c.quiteChan:
			break LOOP
		}

	}
	for {
		if len(c.logDataChan) == 0 {
			break
		}
		c.writelog(<-c.logDataChan)
	}
	c.quiteChan <- struct{}{}
}

//摧毁
func (c *console) Destroy() {
	c.quiteChan <- struct{}{}
	<-c.quiteChan
	close(c.quiteChan)
	close(c.logDataChan)
}

func init() {
	Register(CONSOLE, newConsole)
}
