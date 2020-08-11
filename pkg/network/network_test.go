package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCNAMEWhenItFails(t *testing.T) {
	matches, err := ValidateCNAME("invalid host", "")
	assert.Equal(t, false, matches)
	assert.Equal(t, "lookup invalid host: no such host", err.Error())
}

func TestValidateCNAMEWhenItDoesNotMatch(t *testing.T) {
	matches, err := ValidateCNAME("google.com", "gl.com")
	assert.Equal(t, false, matches)
	assert.Empty(t, err)
}

func TestValidateCNAMEWhenItMatches(t *testing.T) {
	matches, err := ValidateCNAME("google.com", "google.com")
	assert.Equal(t, true, matches)
	assert.Empty(t, err)
}
