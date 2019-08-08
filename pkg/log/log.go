package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func Initialize(format, level string) {
	levelMappings := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}
	logLevel, ok := levelMappings[level]
	if !ok {
		logrus.Warnf("Unknown logging level (%s), falling back to 'info'", level)
		logLevel = logrus.InfoLevel
	}

	if format == "json" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		if format != "text" {
			logrus.Warnf("Unknown logging format (%s), falling back to 'text'", format)
		}
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logLevel)
}
