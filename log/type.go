package log

// LogType 类型
type LogType int

// 日志定义
const (
	LogDebug LogType = iota
	LogInfo
	LogWarn
	LogError
	LogFatal
)

// 日志名称
var names = [...]string{
	LogDebug: "Debug",
	LogInfo:  "Info",
	LogWarn:  "Warn",
	LogError: "Error",
	LogFatal: "Fatal",
}

// 获得类型名称
func (l LogType) String() string {
	return names[l]
}
