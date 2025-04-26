package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	result := AddOne(2)
	expected := 3 // Use 6 if you want this test to fail intentionally

	// Check equality
	assert.Equal(t, expected, result, "they should be equal")

	// Check inequality (just an example)
	assert.NotEqual(t, 5, result, "they should not be equal to 5")

	// Check nil (you need a nil value, here's an example)
	var ptr *int = nil
	assert.Nil(t, ptr, "ptr should be nil")
}
