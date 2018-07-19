// 日志生成器
package alog

type Logger struct {
	ctx *logContext
}

func NewLogger() *Logger {
	return &Logger{
		ctx: context,
	}
}

func (lg *Logger) LogPrint(level int, a ...interface{}) {
	lg.logPrint(level, a)
}

func (lg *Logger) DebugPrint(a  ...interface{}) {
	lg.logPrint(LogLevel_Debug, a)
}

func (lg *Logger) InfoPrint(a  ...interface{}) {
	lg.logPrint(LogLevel_Info, a)
}

func (lg *Logger) NotifyPrint(a  ...interface{}) {
	lg.logPrint(LogLevel_Notify, a)
}

func (lg *Logger) WarnPrint(a  ...interface{}) {
	lg.logPrint(LogLevel_Warn, a)
}

func (lg *Logger) ErrorPrint(a  ...interface{}) {
	lg.logPrint(LogLevel_Error, a)
}

func (lg *Logger) logPrint(level int, args []interface{}) {
	lg.ctx.logPrint(level, args)
}
