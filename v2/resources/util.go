package resources

import (
	"log"
	"os"
	"strings"
)

// HTLogger is a logger which satisfies the go-openapi/runtime/logger.Logger
// interface using go stdlib "log". It appends a prefix to each message to make
// them compatible with what TF expects for legacy logging (see https://developer.hashicorp.com/terraform/plugin/log/writing#legacy-logging).
type HTLogger struct{}

// Printf just prints a log message. The TF "logging level" is hardcoded to
// "DEBUG" since it's expected that this will only be used in the
// go-openapi/runtime to log HTTP transport debugging info.
func (l HTLogger) Printf(format string, args ...interface{}) {
	log.Printf("[DEBUG] HTTP transport: "+format, args...)
}

// Debugf prints a log message the same way as Printf does. The TF "logging
// level" is hardcoded to "DEBUG" since it's expected that this will only
// be used in the go-openapi/runtime to log HTTP transport debugging info.
func (l HTLogger) Debugf(format string, args ...interface{}) {
	log.Printf("[DEBUG] HTTP transport: "+format, args...)
}

// envVarIsEnabled checks if an environment variable is set to a value of
// `yes` or `true` (case insensitive).
func envVarIsEnabled(k string) bool {
	v := strings.TrimSpace(os.Getenv(k))

	return len(v) > 0 && (strings.EqualFold(v, "true") || strings.EqualFold(v, "yes"))
}
