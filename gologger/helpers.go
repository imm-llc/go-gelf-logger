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

	m := wrapBuildGraylogMessage(shortMessage, fullMessage, 6, extraFields, h)
	sendGraylogMessageV2(m)
}

//LogWarning wraps around GrayLogger to make it simpler to log messages
func LogWarning(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) {
	extraFields["tag"] = tag

	h, _ := os.Hostname()

	m := wrapBuildGraylogMessage(shortMessage, fullMessage, 4, extraFields, h)
	sendGraylogMessageV2(m)
}

//LogError wraps around GrayLogger to make it simpler to log messages
func LogError(shortMessage string, fullMessage string, tag string, extraFields map[string]interface{}, err error) {
	extraFields["tag"] = tag
	h, _ := os.Hostname()

	m := wrapBuildGraylogMessage(shortMessage, fullMessage, 3, extraFields, h)
	sendGraylogMessageV2(m)
}
