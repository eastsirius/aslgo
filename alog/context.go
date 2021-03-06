// 全局上下文
package alog

import (
	"time"
	"sync"
)

// 日志上下文
type logContext struct {
	level int
	fmt Formater
	writers []Writer
	lock sync.Locker
	queue chan *LogParam
	quitEvent chan int
}

var context = newLogContext()

func newLogContext() *logContext {
	lc := new(logContext)
	lc.level = LogLevel_Debug
	lc.fmt = NewBasicFormater()
	lc.writers = make([]Writer, 0)
	lc.lock = new(sync.Mutex)
	lc.queue = make(chan *LogParam, 128)
	lc.quitEvent = make(chan int, 1)

	lc.addOutput(NewConsoleWriter())
	go lc.logProc()

	return lc
}

func (ctx *logContext) logPrint(level int, a []interface{}) {
	ctx.logPrintf(level, "", a)
}

func (ctx *logContext) logPrintf(level int, format string, a []interface{}) {
	parm := &LogParam{
		Time: time.Now(),
		Level: level,
		Format: format,
		Args: a,
	}
	ctx.queue <- parm
}

func (ctx *logContext) stop() {
	ctx.queue <- nil
	<- ctx.quitEvent
	ctx.clearOutpus()
}

func (ctx *logContext) setLogLevel(level int) {
	ctx.level = level
}

func (ctx *logContext) setFormater(fmt Formater) {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	ctx.fmt = fmt
}

func (ctx *logContext) addOutput(writer Writer) {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	writer.Start()
	ctx.writers = append(ctx.writers, writer)
}

func (ctx *logContext) clearOutpus() {
	ctx.lock.Lock()
	defer ctx.lock.Unlock()

	for _, v := range ctx.writers {
		v.Stop()
	}
	ctx.writers = make([]Writer, 0)
}

func (ctx *logContext) logProc() {
	for true {
		parm := <- ctx.queue

		if parm == nil {
			ctx.quitEvent <- 0
			return
		}

		if parm.Level < ctx.level {
			continue
		}

		ctx.lock.Lock()
		formater := ctx.fmt
		writers := ctx.writers
		ctx.lock.Unlock()

		log := &LogItem{
			Time: parm.Time,
			Level: parm.Level,
			Log: formater.Format(parm),
		}
		for _, v := range writers {
			v.WriteLog(log)
		}
	}
}
