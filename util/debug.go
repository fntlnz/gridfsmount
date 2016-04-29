package util

import (
	"os"

	"github.com/Sirupsen/logrus"
)

const DEBUG_ENVVAR = "DEBUG"

func EnableDebug() {
	os.Setenv(DEBUG_ENVVAR, "1")
	logrus.SetLevel(logrus.DebugLevel)
}

func DisableDebug() {
	os.Unsetenv(DEBUG_ENVVAR)
	logrus.SetLevel(logrus.InfoLevel)
}

func IsDebugEnabled() bool {
	if value, ok := os.LookupEnv(DEBUG_ENVVAR); ok && value == "1" {
		return true
	}
	return false
}
