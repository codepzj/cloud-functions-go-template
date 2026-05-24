package database

import (
	glogger "gorm.io/gorm/logger"
)

func parseLogLevel(level string) glogger.LogLevel {
	switch level {
	case "silent":
		return glogger.Silent
	case "warn":
		return glogger.Warn
	case "error":
		return glogger.Error
	}
	return glogger.Info
}
