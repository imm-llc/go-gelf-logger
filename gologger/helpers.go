package gologger

func buildGraylogLogMessage(s string, f string, ef map[string]interface{}, e error) LogItems {
	return LogItems{
		ShortMsg:    s,
		FullMsg:     f,
		ExtraFields: ef,
		ErrorMsg:    e,
	}
}

//LogInfo wraps around GrayLogger to make it simpler to log messages.
func LogInfo(shortMessage string, fullMessage string, extraFields map[string]interface{}, err error) error {
	return GrayLogger.Info(buildGraylogLogMessage(shortMessage, fullMessage, extraFields, err))
}

func LogWarning(shortMessage string, fullMessage string, extraFields map[string]interface{}, err error) error {
	return GrayLogger.Warning(buildGraylogLogMessage(shortMessage, fullMessage, extraFields, err))
}

func LogError(shortMessage string, fullMessage string, extraFields map[string]interface{}, err error) error {
	return GrayLogger.Error(buildGraylogLogMessage(shortMessage, fullMessage, extraFields, err))
}
