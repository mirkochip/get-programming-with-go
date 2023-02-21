package experiments

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMalacandra(t *testing.T) {
	expected := 100.0
	result := speedCalculator(500.0, 5.0)
	assert.Equal(t, expected, result, "The two speeds should be the same.")
}
