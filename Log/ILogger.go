package Log

type ILogger interface {
	LogInfo(message string)
	LogWarn(message string)
	LogError(err error)
}
