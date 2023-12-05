package go_mesiab_sdk

import (
	"bytes"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// TestGetLogger tests the singleton behavior of GetLogger.
func TestGetLogger(t *testing.T) {
	logger1 := GetLogger()
	logger2 := GetLogger()

	assert.Equal(t, logger1, logger2, "Expected GetLogger to return the same instance")
}

// TestLogMessageFormatting tests the formatting of log messages.
func TestLogMessageFormatting(t *testing.T) {
	msg := Logf("test %s", "message")

	assert.Equal(t, "test message", msg.Message, "Logf should format message correctly")
}

// TestLogMessageAdd tests adding fields to a log message.
func TestLogMessageAdd(t *testing.T) {
	msg := Logf("test").Add("key", "value")

	assert.Equal(t, "value", msg.Fields["key"], "Add should insert the correct value for a key")
}

// TestLogOutput tests the output of the logger at different levels.
func TestLogOutput(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	msg := "test message"
	Logf(msg).Info()

	assert.Contains(t, buf.String(), msg, "Info should log the correct message")

	buf.Reset()
	Logf(msg).Debug()
	assert.Contains(t, buf.String(), "🐛 "+msg, "Debug should log the correct message with emoji")
}

// TestAddError tests adding an error and stack trace to the log.
func TestAddError(t *testing.T) {
	err := assert.AnError
	msg := Logf("error").AddError(err)

	assert.Equal(t, err.Error(), msg.Fields["error"], "AddError should add the correct error message")
	assert.NotEmpty(t, msg.Fields["stack"], "AddError should add a stack trace")
}
