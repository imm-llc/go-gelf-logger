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
func LogInfo(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) error {
	extraFields["tag"] = tag
	return GrayLogger.Info(buildGraylogLogMessage(shortMessage, fullMessage, extraFields, err))
}

//LogWarning wraps around GrayLogger to make it simpler to log messages
func LogWarning(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) error {
	extraFields["tag"] = tag
	return GrayLogger.Warning(buildGraylogLogMessage(shortMessage, fullMessage, extraFields, err))
}

//LogError wraps around GrayLogger to make it simpler to log messages
func LogError(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) error {
	extraFields["tag"] = tag
	return GrayLogger.Error(buildGraylogLogMessage(shortMessage, fullMessage, extraFields, err))
}
