package rerr

const (
	UserMsgKey         = "userMsgKey"         // user message key
	UserMsgSeverityKey = "userMsgSeverityKey" // user message severity key
)

var Severity severity

type severity struct {
	Success, Error, Warn, Info string
}

func init() {
	Severity = severity{
		Success: "success",
		Error:   "error",
		Warn:    "info",
		Info:    "warn",
	}
}

type UserMsgOptions struct {
	Severity string
}
