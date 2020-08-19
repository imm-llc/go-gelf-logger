package gologger

import "os"

func buildGraylogLogMessage(s string, f string, ef map[string]interface{}, e error) LogItems {
	return LogItems{
		ShortMsg:    s,
		FullMsg:     f,
		ExtraFields: ef,
		ErrorMsg:    e,
	}
}

//LogInfo wraps around GrayLogger to make it simpler to log messages.
func LogInfo(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) {
	extraFields["tag"] = tag
	h, _ := os.Hostname()

	sendGraylogMessage(wrapBuildGraylogMessage(shortMessage, fullMessage, 6, extraFields, h))
}

func LogWarning(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) {
	extraFields["tag"] = tag
	h, _ := os.Hostname()
	sendGraylogMessage(wrapBuildGraylogMessage(shortMessage, fullMessage, 4, extraFields, h))
}

func LogError(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) {
	extraFields["tag"] = tag
	h, _ := os.Hostname()
	sendGraylogMessage(wrapBuildGraylogMessage(shortMessage, fullMessage, 3, extraFields, h))
}
