package log

// LogPort 日志接口
type LogPort interface {
	// 日志记录
	LogError(format string, v ...interface{})
	LogWarn(format string, v ...interface{})
	LogInfo(format string, v ...interface{})
	LogDebug(format string, v ...interface{})
}
