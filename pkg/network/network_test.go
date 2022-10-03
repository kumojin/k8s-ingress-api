package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCNAMEWhenItFails(t *testing.T) {
	ok, err := ValidateCNAME("invalid host", "")
	assert.Equal(t, false, ok)
	assert.Equal(t, "lookup invalid host: no such host", err.Error())
}

func TestValidateCNAMEWhenItDoesNotMatch(t *testing.T) {
	ok, err := ValidateCNAME("m.facebook.com", "star-mini.facebook.com")
	assert.Equal(t, false, ok)
	assert.Empty(t, err)
}

func TestValidateCNAMEWhenItMatches(t *testing.T) {
	ok, err := ValidateCNAME("m.facebook.com", "star-mini.c10r.facebook.com")
	assert.Equal(t, true, ok)
	assert.Empty(t, err)
}
