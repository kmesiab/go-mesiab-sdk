package go_mesiab_sdk

import (
	"os"
	"sync"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetConfigSingleton tests if GetConfig returns the same instance every time.
func TestGetConfigSingleton(t *testing.T) {
	oldLogLevel := os.Getenv("LOG_LEVEL")
	err := os.Setenv("LOG_LEVEL", "info")
	if err != nil {
		log.Errorf("Failed to set LOG_LEVEL environment variable")
	}

	defer func() {
		err := os.Setenv("LOG_LEVEL", oldLogLevel)
		if err != nil {
			log.Errorf("Failed to reset LOG_LEVEL environment variable")
		}
	}()

	// Reset global variables for a clean test environment
	config = nil
	once = sync.Once{}

	firstConfig, err := GetConfig()
	assert.Nil(t, err, "GetConfig should not return an error")
	assert.NotNil(t, firstConfig, "GetConfig should return a Config instance")

	secondConfig, _ := GetConfig()
	assert.Equal(t, firstConfig, secondConfig, "GetConfig should return the same instance")
}

// TestValidateConfig tests the validation of the Config struct.
func TestValidateConfig(t *testing.T) {
	oldLogLevel := os.Getenv("LOG_LEVEL")
	err := os.Setenv("LOG_LEVEL", "info")
	if err != nil {
		log.Errorf("Failed to set LOG_LEVEL environment variable")
	}

	defer func() {
		err := os.Setenv("LOG_LEVEL", oldLogLevel)
		if err != nil {
			log.Errorf("Failed to reset LOG_LEVEL environment variable")
		}
	}()

	// Test with a valid Config
	validConfig, err := GetConfig()
	require.NoError(t, err, "GetConfig should not return an error")

	err = ValidateConfig(validConfig)
	assert.Nil(t, err, "ValidateConfig should not return an error for a valid config")

	// Test with a nil Config
	err = ValidateConfig(nil)
	assert.NotNil(t, err, "ValidateConfig should return an error for a nil config")
}
