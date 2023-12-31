// Package go_mesiab_sdk Package provides a logging utility designed for AWS Lambda functions
// and API Gateway. It offers a globally accessible logger configured to output
// JSON formatted logs. The package also includes utility functions for adding
// additional metadata fields to logs, either from custom fields or directly
// from AWS API Gateway and Custom Authorizer request objects.
//
// Example Usage:
//
//	// Initializing and getting the logger
//	logger := GetLogger()
//
//	// Logging with a simple message
//	PrepareLogMessage("A simple message").Info()
//
//	// Logging with additional fields
//	PrepareLogMessage("With extra fields").Add("field1", "value1").Warn()
//
//	// Logging with API Gateway request details
//	PrepareLogMessage("API Gateway request info").AddGatewayRequest(request).Info()
//
// Global Logger:
//
// The logger is initialized only once and is globally accessible via GetLogger().
// It's configured to output JSON-formatted logs and includes metadata like 'env'
// and 'app_name' from environment variables.
//
// Custom Log Messages:
//
// The LogMessage struct provides a way to prepare a log message with additional
// metadata fields. You can add custom fields using the Add() method or append
// AWS request details using AddAuthorizerRequest() and AddGatewayRequest().
package go_mesiab_sdk

import (
	"fmt"
	"os"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

var globalLogger *log.Entry // Singleton logger instance

// GetLogger initializes and returns the global logger.
func GetLogger() *log.Entry {
	if globalLogger == nil {
		// Logger setup
		log.SetFormatter(&log.TextFormatter{
			ForceColors:      true,
			DisableTimestamp: true,
		})
		log.SetReportCaller(false)
		log.SetOutput(os.Stdout)
		globalLogger = log.WithFields(log.Fields{
			"env":      os.Getenv("ENV"),
			"app_name": os.Getenv("APP_NAME"),
		})
	}

	return globalLogger
}

// Logf creates a new LogMessage with a formatted message.
func Logf(format string, args ...interface{}) *LogMessage {
	return &LogMessage{Message: fmt.Sprintf(format, args...), Fields: make(map[string]interface{})}
}

// LogMessage holds the log message and additional fields.
type LogMessage struct {
	Message string     `json:"message"`
	Fields  log.Fields `json:"fields"`
}

// Add adds a key-value pair to the LogMessage's Fields
// Chainable: Can be chained with other methods.
func (l LogMessage) Add(key string, value string) LogMessage {
	l.Fields[key] = value

	return l
}

func (l LogMessage) AddError(err error) LogMessage {
	trace := debug.Stack()

	l.Fields["error"] = err.Error()
	l.Fields["stack"] = fmt.Sprintf("%+v", trace)

	return l
}

// Info logs the message at Info level
// Chainable: Can be chained with other methods.
func (l LogMessage) Info() {
	GetLogger().WithFields(l.Fields).Info(l.Message)
}

// Debug logs the message at Debug level
// Chainable: Can be chained with other methods.
func (l LogMessage) Debug() {
	GetLogger().WithFields(l.Fields).Debug("🐛 " + l.Message)
}

// Warn logs the message at Warn level
// Chainable: Can be chained with other methods.
func (l LogMessage) Warn() {
	GetLogger().WithFields(l.Fields).Warn("⚠️ " + l.Message)
}

// Error logs the message at Error level
// Chainable: Can be chained with other methods.
func (l LogMessage) Error() {
	GetLogger().WithFields(l.Fields).Error(l.Message)
}

// Fatal logs the message at Fatal level
// Chainable: Can be chained with other methods.
func (l LogMessage) Fatal() {
	GetLogger().WithFields(l.Fields).Fatal(l.Message)
}
