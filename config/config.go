package config

type Log struct {
	Level string
	Formatter string
	ReportCaller bool
}

var LogSetting = &Log{}


func loadLog
