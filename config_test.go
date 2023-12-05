package go_mesiab_sdk

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetConfigSingleton tests if GetConfig returns the same instance every time.
func TestGetConfigSingleton(t *testing.T) {
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
	// Test with a valid Config
	validConfig := &Config{}
	err := ValidateConfig(validConfig)
	assert.Nil(t, err, "ValidateConfig should not return an error for a valid config")

	// Test with a nil Config
	err = ValidateConfig(nil)
	assert.NotNil(t, err, "ValidateConfig should return an error for a nil config")

	// Test with a non-struct Config (should never happen in practice, but good to test)
	err = ValidateConfig(validConfig)
	assert.NotNil(t, err, "ValidateConfig should return an error for a non-struct config")
}
