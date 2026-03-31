package models

type Severity string

const (
	SeverityError Severity = "ERROR"
	SeverityWarn  Severity = "WARN"
	SeverityInfo  Severity = "INFO"
)

type Finding struct {
	Severity Severity
	Resource string
	Message  string
	Rule     string
	File     string
}
