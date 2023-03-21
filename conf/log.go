package conf

// LogFormat 日志格式
type LogFormat string

// LogTo 记录到哪儿
type LogTo string

const (
	TextFormat = LogFormat("text")
	JSONFormat = LogFormat("json")
)

const (
	ToFile   = LogTo("file")
	ToStdout = LogTo("ToStdout")
)
