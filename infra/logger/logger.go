package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// SetupLogger configures the logrus instance
func SetupLogger() {
	if os.Getenv("LOGGER_DEV") == "true" {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	logrus.SetOutput(os.Stdout)
}
