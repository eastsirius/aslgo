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

func (lg *Logger) DebugPrint(a ...interface{}) {
	lg.logPrint(LogLevel_Debug, a)
}

func (lg *Logger) InfoPrint(a ...interface{}) {
	lg.logPrint(LogLevel_Info, a)
}

func (lg *Logger) NotifyPrint(a ...interface{}) {
	lg.logPrint(LogLevel_Notify, a)
}

func (lg *Logger) WarnPrint(a ...interface{}) {
	lg.logPrint(LogLevel_Warn, a)
}

func (lg *Logger) ErrorPrint(a ...interface{}) {
	lg.logPrint(LogLevel_Error, a)
}

func (lg *Logger) LogPrintf(level int, format string, a ...interface{}) {
	lg.logPrintf(level, format, a)
}

func (lg *Logger) DebugPrintf(format string, a ...interface{}) {
	lg.logPrintf(LogLevel_Debug, format, a)
}

func (lg *Logger) InfoPrintf(format string, a ...interface{}) {
	lg.logPrintf(LogLevel_Info, format, a)
}

func (lg *Logger) NotifyPrintf(format string, a ...interface{}) {
	lg.logPrintf(LogLevel_Notify, format, a)
}

func (lg *Logger) WarnPrintf(format string, a ...interface{}) {
	lg.logPrintf(LogLevel_Warn, format, a)
}

func (lg *Logger) ErrorPrintf(format string, a ...interface{}) {
	lg.logPrintf(LogLevel_Error, format, a)
}

func (lg *Logger) logPrint(level int, args []interface{}) {
	lg.ctx.logPrint(level, args)
}

func (lg *Logger) logPrintf(level int, format string, args []interface{}) {
	lg.ctx.logPrintf(level, format, args)
}
